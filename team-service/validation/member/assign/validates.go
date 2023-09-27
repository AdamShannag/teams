package assign

import (
	"context"
	"errors"
	"fmt"
	"team-service/repository/ent/member"
	"team-service/repository/ent/team"
	resource "team-service/resource/member"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validator) validateTeamId(teamId string, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmptyString(teamId); err != nil {
		return []violation.Violation{violation.FieldViolation("teamId", err)}
	}
	if err := v.existById(teamId, ctx); err != nil {
		return []violation.Violation{violation.FieldViolation("teamId", err)}
	}
	return violations
}

func (v *Validator) validateMembers(request resource.AssignRequest, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmpty(request.Members); err != nil {
		return []violation.Violation{violation.FieldViolation("members", err)}
	}
	for _, memberId := range request.Members {
		if ok := common.IsExistMember(memberId, *v.client, ctx); !ok {
			violations = append(violations, violation.FieldViolation("members", fmt.Errorf("member [%s] not found", memberId)))
		}
		if ok, vo := v.checkMemberAssignability(request.IsAssign(), memberId, ctx); !ok {
			violations = append(violations, vo)
		}
	}
	return violations
}

func (v *Validator) checkMemberAssignability(isAssign bool, memberID string, ctx context.Context) (ok bool, vio violation.Violation) {
	isAssignable := v.assignableMember(memberID, ctx)

	if isAssign {
		if !isAssignable {
			return false, violation.FieldViolation("members", fmt.Errorf("member [%s] is already assigned to a team", memberID))
		}
	} else if isAssignable {
		return false, violation.FieldViolation("members", fmt.Errorf("member [%s] is not assigned to a team", memberID))
	}
	return true, vio
}

func (v *Validator) existById(id string, ctx context.Context) error {
	if ok, _ := v.client.Team.
		Query().
		Where(team.ID(id)).
		Exist(ctx); !ok {
		return errors.New(fmt.Sprintf("team id [%s] dose not exist", id))
	}
	return nil
}

func (v *Validator) assignableMember(memberId string, ctx context.Context) (ok bool) {
	ok, _ = v.client.Member.
		Query().
		Where(member.StatusEQ(member.StatusFREE)).
		Where(member.ID(memberId)).
		Exist(ctx)
	return
}
