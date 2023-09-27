package approval

import (
	"context"
	"fmt"
	"team-service/repository/ent/member"
	"team-service/validation/common"
	"team-service/validation/violation"
)

func (v *Validator) validateAssignApproval(memberIds []string, ctx context.Context) (violations []violation.Violation) {
	if err := common.IsEmpty(memberIds); err != nil {
		return []violation.Violation{violation.FieldViolation("members", err)}
	}
	for _, memberId := range memberIds {
		if ok := common.IsExistMember(memberId, *v.client, ctx); !ok {
			violations = append(violations, violation.FieldViolation("members", fmt.Errorf("member [%s] not found", memberId)))
		}
		if ok := v.isPendingMember(memberId, ctx); !ok {
			violations = append(violations, violation.FieldViolation("members", fmt.Errorf("member [%s] is not pending assignation", memberId)))
		}
	}
	return violations
}

func (v *Validator) isPendingMember(memberId string, ctx context.Context) (ok bool) {
	ok, _ = v.client.Member.
		Query().
		Where(member.StatusEQ(member.StatusPENDING)).
		Where(member.ID(memberId)).
		Exist(ctx)
	return
}
