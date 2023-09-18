package userservice

import (
	"context"
	"errors"
	"user-service/model"

	"github.com/Nerzal/gocloak/v13"
)

func (k *keycloakUserServiceImpl) Create(ctx context.Context, u *model.CreateUserRequest) (string, error) {
	token, err := k.token(ctx)

	if err != nil {
		return "", errors.New("unkown client credentials")
	}

	user := gocloak.User{
		FirstName:     gocloak.StringP(u.FirstName),
		LastName:      gocloak.StringP(u.LastName),
		Email:         gocloak.StringP(u.Email),
		Enabled:       gocloak.BoolP(true),
		Username:      gocloak.StringP(u.Username),
		EmailVerified: gocloak.BoolP(false),
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Type:      gocloak.StringP("password"),
				Value:     gocloak.StringP("P@ssw0rd"),
				Temporary: gocloak.BoolP(true),
			},
		},
	}

	userID, err := k.client.CreateUser(ctx, token.AccessToken, k.realm, user)
	if err != nil {
		return "", errors.New("failed to create user")
	}

	return userID, nil
}
