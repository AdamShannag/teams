package users

import (
	"errors"
	"net/http"
	"user-service/cmd/user-service/config"
	"user-service/model"
	"user-service/pkg/nts"

	"github.com/AdamShannag/toolkit/v2"
	"github.com/nats-io/nats.go"
)

func (u *Users) Create(w http.ResponseWriter, req *http.Request) {
	var createRequest model.CreateUserRequest
	err := u.ReadJSON(w, req, &createRequest)
	if err != nil {
		u.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	userID, err := u.keycloakUserService.Create(req.Context(), &createRequest)

	if err != nil {
		u.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_, err = nts.Publish(req.Context(), u, &nats.Msg{
		Subject: config.USERS_SUBJECT_NEW,
		Data:    []byte(userID),
	})

	if err != nil {
		u.Error().Err(err).Msg("an error has occurred while publish")
		u.ErrorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}

	if err := u.WriteJSON(w, http.StatusCreated, toolkit.JSONResponse{
		Error:   false,
		Message: "user created successfully",
	}); err != nil {
		u.Error().Err(err).Msg("an error has occurred")
		u.ErrorJSON(w, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}
}
