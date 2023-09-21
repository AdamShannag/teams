package koj

import (
	"context"
	"crypto/rsa"
	"strings"
	"user-service/pkg/logger"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
)

type KeycloakOfflineJWT struct {
	zerolog.Logger
	publicKey *rsa.PublicKey
	mode      KeyclaokMode
	client    *gocloak.GoCloak
	relam     string
}

type KeyclaokMode string

const (
	PEM_HEADER string = "-----BEGIN PUBLIC KEY-----\n"
	PEM_FOOTER string = "\n-----END PUBLIC KEY-----"
)

const (
	KEYCLOAK_PROD KeyclaokMode = "prod"
	KEYCLOAK_DEV  KeyclaokMode = "dev"
)

func NewKeycloakOfflineJWT(client *gocloak.GoCloak, realm string, mode KeyclaokMode) *KeycloakOfflineJWT {
	kj := &KeycloakOfflineJWT{
		Logger: logger.Get(),
		mode:   mode,
		client: client,
		relam:  realm,
	}
	kj.getPublicKey()
	return kj
}

func (k *KeycloakOfflineJWT) getPublicKey() {
	if k.mode == KEYCLOAK_DEV {
		return
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(k.getPublickKeyFromClient())
	k.fatal(err, "error parsing public key")

	k.publicKey = key
}

func (k *KeycloakOfflineJWT) getPublickKeyFromClient() []byte {
	cr, err := k.client.GetCerts(context.Background(), k.relam)
	k.fatal(err, "error obtaining public key from auth server")

	x5cs := (*cr.Keys)[1].X5c
	rawKey := (*x5cs)[0]

	sb := strings.Builder{}

	sb.WriteString(PEM_HEADER)
	sb.WriteString(rawKey)
	sb.WriteString(PEM_FOOTER)

	return []byte(sb.String())

}

func (k *KeycloakOfflineJWT) ValidateToken(ts string) (*KeycloakAccessToken, error) {
	if k.mode == KEYCLOAK_DEV {
		return nil, nil
	}
	token, err := jwt.ParseWithClaims(ts, &KeycloakAccessToken{}, func(token *jwt.Token) (interface{}, error) {
		return k.publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*KeycloakAccessToken)

	return claims, nil
}

func (k *KeycloakOfflineJWT) fatal(err error, msg string) {
	if err != nil {
		k.Fatal().Err(err).Msg(msg)
	}
}
