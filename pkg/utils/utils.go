package utils

import (
	"strings"

	"github.com/influenzanet/go-utils/pkg/api_types"
)

// IsTokenEmpty check a token from api if it's empty
func IsTokenEmpty(t *api_types.TokenInfos) bool {
	if t == nil || t.Id == "" || t.InstanceId == "" {
		return true
	}
	return false
}

func GetUsernameFromToken(t *api_types.TokenInfos) string {
	if t == nil {
		return ""
	}
	username, ok := t.Payload["username"]
	if !ok {
		return ""
	}
	return username
}

// CheckRoleInToken Check if role is present in the token
func CheckRoleInToken(t *api_types.TokenInfos, role string) bool {
	if t == nil {
		return false
	}
	if val, ok := t.Payload["roles"]; ok {
		roles := strings.Split(val, ",")
		for _, r := range roles {
			if r == role {
				return true
			}
		}
	}
	return false
}

// CheckIfAnyRolesInToken if token contains any of the roles
func CheckIfAnyRolesInToken(t *api_types.TokenInfos, requiredRoles []string) bool {
	if t == nil {
		return false
	}
	if val, ok := t.Payload["roles"]; ok {
		roles := strings.Split(val, ",")
		for _, r := range roles {
			for _, rr := range requiredRoles {
				if r == rr {
					return true
				}
			}
		}
	}
	return false
}
