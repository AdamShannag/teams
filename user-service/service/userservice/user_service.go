package userservice

import (
	"context"
	"user-service/model"

	"github.com/Nerzal/gocloak/v13"
)

type KeycloakUserService interface {
	Create(context.Context, *model.CreateUserRequest) (string, error)
	token(ctx context.Context) (*gocloak.JWT, error)
}

type keycloakUserServiceImpl struct {
	client                        *gocloak.GoCloak
	clientId, clientSecret, realm string
}

func NewKeycloakUserServiceImpl(client *gocloak.GoCloak, clientId, clientSecret, realm string) *keycloakUserServiceImpl {
	return &keycloakUserServiceImpl{
		client,
		clientId,
		clientSecret,
		realm,
	}
}
