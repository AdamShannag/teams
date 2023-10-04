package assign

import (
	"context"
	"errors"
	"fmt"
	resource "project-service/resource/member"
	"team-service/repository/ent/member"
	"team-service/repository/ent/team"
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
		if ok := v.isExistMember(memberId, ctx); !ok {
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
			return false, violation.FieldViolation("members", fmt.Errorf("member [%s] is already assigned to a project", memberID))
		}
	} else if isAssignable {
		return false, violation.FieldViolation("members", fmt.Errorf("member [%s] is not assigned to a project", memberID))
	}
	return true, vio
}

func (v *Validator) existById(id string, ctx context.Context) error {
	ok, err := v.teamRepo.ExistByIdAndStatusNot(ctx, id, team.StatusDELETED)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(fmt.Sprintf("project id [%s] dose not exist", id))
	}
	return nil
}

func (v *Validator) assignableMember(memberId string, ctx context.Context) (ok bool) {
	ok, _ = v.memberRepo.ExistByIdAndStatus(ctx, memberId, member.StatusFREE)
	return
}

func (v *Validator) isExistMember(memberId string, ctx context.Context) (ok bool) {
	ok, _ = v.memberRepo.ExistById(ctx, memberId)
	return
}
