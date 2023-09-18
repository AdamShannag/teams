package userservice

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
)

func (k *keycloakUserServiceImpl) token(ctx context.Context) (*gocloak.JWT, error) {
	return k.client.LoginClient(
		ctx,
		k.clientId,
		k.clientSecret,
		k.realm,
	)
}
