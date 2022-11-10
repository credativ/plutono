package teamguardian

import (
	"github.com/credativ/plutono/pkg/bus"
	"github.com/credativ/plutono/pkg/models"
)

func CanAdmin(bus bus.Bus, orgId int64, teamId int64, user *models.SignedInUser) error {
	if user.OrgRole == models.ROLE_ADMIN {
		return nil
	}

	if user.OrgId != orgId {
		return models.ErrNotAllowedToUpdateTeamInDifferentOrg
	}

	cmd := models.GetTeamMembersQuery{
		OrgId:  orgId,
		TeamId: teamId,
		UserId: user.UserId,
	}

	if err := bus.Dispatch(&cmd); err != nil {
		return err
	}

	for _, member := range cmd.Result {
		if member.UserId == user.UserId && member.Permission == models.PERMISSION_ADMIN {
			return nil
		}
	}

	return models.ErrNotAllowedToUpdateTeam
}
