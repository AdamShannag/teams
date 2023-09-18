package main

import (
	"user-service/cmd/user-service/config"

	"github.com/Nerzal/gocloak/v13"
)

func getKeycloakClient() *gocloak.GoCloak {
	return gocloak.NewClient(config.KEYCLOAK_BASE_URL)
}
