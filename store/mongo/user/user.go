package user

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
)

const (
	EntityType    = "user"
	StatusActive  = "active"
	StatusBlocked = "blocked"
)

func AnyRole4DomainSiteAppId(roles []commons.UserRole, domain, site, appId, appType, role string) bool {
	return model.AnyRole4DomainSiteAppId(roles, domain, site, appId, appType, role)
}
