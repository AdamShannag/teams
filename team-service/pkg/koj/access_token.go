package koj

import "github.com/golang-jwt/jwt/v5"

type KeycloakAccessToken struct {
	Exp            *jwt.NumericDate `json:"exp"`
	Iat            *jwt.NumericDate `json:"iat"`
	AuthTime       int              `json:"auth_time"`
	Jti            string           `json:"jti"`
	Iss            string           `json:"iss"`
	Aud            jwt.ClaimStrings `json:"aud"`
	Sub            string           `json:"sub"`
	Typ            string           `json:"typ"`
	Azp            string           `json:"azp"`
	Nonce          string           `json:"nonce"`
	SessionState   string           `json:"session_state"`
	Acr            string           `json:"acr"`
	AllowedOrigins []string         `json:"allowed-origins"`
	RealmAccess    RealmAccess      `json:"realm_access"`
	ResourceAccess struct {
		Account struct {
			Roles []string `json:"roles"`
		} `json:"account"`
	} `json:"resource_access"`
	Scope             string `json:"scope"`
	Sid               string `json:"sid"`
	EmailVerified     bool   `json:"email_verified"`
	Name              string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	Email             string `json:"email"`
}

type RealmAccess struct {
	Roles []string `json:"roles"`
}

func (c KeycloakAccessToken) GetExpirationTime() (*jwt.NumericDate, error) {
	return c.Exp, nil
}

func (c KeycloakAccessToken) GetNotBefore() (*jwt.NumericDate, error) {
	return nil, nil
}

func (c KeycloakAccessToken) GetIssuedAt() (*jwt.NumericDate, error) {
	return c.Iat, nil
}

func (c KeycloakAccessToken) GetAudience() (jwt.ClaimStrings, error) {
	return c.Aud, nil
}

func (c KeycloakAccessToken) GetIssuer() (string, error) {
	return c.Iss, nil
}

func (c KeycloakAccessToken) GetSubject() (string, error) {
	return c.Sub, nil
}
