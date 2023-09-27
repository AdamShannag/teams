package common

import (
	"context"
	"errors"
	"team-service/repository/ent"
	"team-service/repository/ent/member"
)

func IsNilString(value *string) error {
	if value == nil {
		return errors.New("is required")
	}
	return nil
}

func IsEmptyString(value string) error {
	if len(value) == 0 {
		return errors.New("is required")
	}
	return nil
}

func IsEmpty(value []string) error {
	if value == nil || len(value) == 0 {
		return errors.New("is required")
	}
	return nil
}

func IsExistMember(memberId string, client ent.Client, ctx context.Context) (ok bool) {
	ok, _ = client.Member.
		Query().
		Where(member.ID(memberId)).
		Exist(ctx)
	return
}
