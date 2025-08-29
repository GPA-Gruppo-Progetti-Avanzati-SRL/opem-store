package commons

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

// @tpm-schematics:start-region("top-file-section")
// @tpm-schematics:end-region("top-file-section")

type UserRole struct {
	Domain string `json:"domain,omitempty" bson:"domain,omitempty" yaml:"domain,omitempty"`
	Site   string `json:"site,omitempty" bson:"site,omitempty" yaml:"site,omitempty"`
	Apps   string `json:"apps,omitempty" bson:"apps,omitempty" yaml:"apps,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s UserRole) IsZero() bool {
	return s.Domain == "" && s.Site == "" && s.Apps == ""
}

// @tpm-schematics:start-region("bottom-file-section")

func (rd UserRole) MatchDomainAndSite(domain, site string) bool {
	if rd.Domain != "*" && domain != rd.Domain {
		return false
	}

	if rd.Site != "*" && site != rd.Site {
		return false
	}

	return true
}

type UserRoleList []UserRole

func (roles UserRoleList) Add(domain, site, appId, appType, appRole string) (UserRoleList, error) {
	const semLogContext = "user::set-role"

	newList := make(UserRoleList, len(roles))
	copy(newList, roles)

	foundRoleNdx := -1
	for i, role := range newList {
		if role.MatchDomainAndSite(domain, site) {
			foundRoleNdx = i
			break
		}
	}

	if foundRoleNdx < 0 {
		newList = append(newList, UserRole{
			Domain: domain,
			Site:   site,
			Apps:   fmt.Sprintf("%s:%s:%s", appId, appType, appRole),
		})
	} else {
		appRoleSet, err := newList[foundRoleNdx].ParseApps()
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
		}

		if appRoleSet.IsZero() {
			newList = append(newList, UserRole{
				Domain: domain,
				Site:   site,
				Apps:   fmt.Sprintf("%s:%s:%s", appId, appType, appRole),
			})
		} else {
			mt, mndx := appRoleSet.MatchTypeAndIndex(appId, appType, appRole)
			switch mt {
			case MatchTypeNone:
				appRoleSet = append(appRoleSet, AppRoleDefinition{
					AppId:    appId,
					AppType:  appType,
					AppRoles: appRole,
				})
			case MatchTypeAppId:
				appRoleSet[mndx].AppType = "*"
				appRoleSet[mndx].AppRoles = appRole
			case MatchTypeAppType:
				appRoleSet[mndx].AppRoles = appRole
			case MatchTypeRole:
			default:
				log.Error().Int("match-type", mt).Msg(semLogContext + " - unexpected app role match type")
			}

			newList[foundRoleNdx].Apps = appRoleSet.String()
		}
	}

	return newList, nil
}

func (roles UserRoleList) Del(domain, site, appId, appType, appRole string) (UserRoleList, error) {
	const semLogContext = "user::del-role"

	newList := make(UserRoleList, len(roles))
	copy(newList, roles)

	foundRoleNdx := -1
	for i, role := range newList {
		if role.MatchDomainAndSite(domain, site) {
			foundRoleNdx = i
			break
		}
	}

	if foundRoleNdx < 0 {
		log.Warn().Str("domain", domain).Str("site", site).Str("app-id", appId).Str("app-type", appType).Str("app-role", appRole).Msg(semLogContext + " - roles not found")
		return newList, nil
	}

	appRoleSet, err := newList[foundRoleNdx].ParseApps()
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
	}

	if appRoleSet.IsZero() {
		log.Warn().Str("domain", domain).Str("site", site).Str("app-id", appId).Str("app-type", appType).Str("app-role", appRole).Msg(semLogContext + " - roles not found")
		return newList, nil
	}

	// app-home:admin:role-requested
	// *:*:admin
	mt, mndx := appRoleSet.MatchTypeAndIndex(appId, appType, appRole)
	switch mt {
	case MatchTypeAppId:
		fallthrough
	case MatchTypeAppType:
		fallthrough
	case MatchTypeNone:
		log.Warn().Str("domain", domain).Str("site", site).Str("app-id", appId).Str("app-type", appType).Str("app-role", appRole).Msg(semLogContext + " - roles not found")
	case MatchTypeRole:
		appRoleSet = appRoleSet.Slice(mndx)
		if len(appRoleSet) > 0 {
			newList[foundRoleNdx].Apps = appRoleSet.String()
		}
	default:
		log.Error().Int("match-type", mt).Msg(semLogContext + " - unexpected app role match type")
	}

	return newList, nil
}

type AppRoleDefinition struct {
	AppId    string
	AppType  string
	AppRoles string
}

const (
	MatchTypeNone = iota
	MatchTypeAppId
	MatchTypeAppType
	MatchTypeRole
)

//func (rd AppRoleDefinition) MatchDomainAndSite(domain, site string) bool {
//	if rd.domain != "*" && domain != rd.domain {
//		return false
//	}
//
//	if rd.site != "*" && site != rd.site {
//		return false
//	}
//
//	return true
//}

func (rd AppRoleDefinition) MatchType(appId, appType, role string) int {

	if rd.AppId != "*" && appId != rd.AppId {
		return MatchTypeNone
	}

	if rd.AppType != "*" && appType != rd.AppType {
		return MatchTypeAppId
	}

	if rd.AppRoles != "*" && role != rd.AppRoles {
		return MatchTypeAppType
	}

	return MatchTypeRole
}

type AppRoleDefinitionSet []AppRoleDefinition

func (s UserRole) ParseApps() (AppRoleDefinitionSet, error) {
	const semLogContext = "user-role::parse"
	var result AppRoleDefinitionSet
	if s.Apps == "" {
		return result, nil
	}

	apps := strings.Split(s.Apps, ";")
	for _, app := range apps {
		appDefParts := strings.Split(app, ":")

		appDef := AppRoleDefinition{
			AppId: appDefParts[0],
		}

		switch len(appDefParts) {
		case 1:
			appDef.AppType = "*"
			appDef.AppRoles = "*"
		case 2:
			appDef.AppType = appDefParts[1]
			appDef.AppRoles = "*"
		case 3:
			appDef.AppType = appDefParts[1]
			appDef.AppRoles = appDefParts[2]
		default:
			err := fmt.Errorf("invalid app definition: %s", app)
			log.Error().Err(err).Msg(semLogContext)
			return nil, err
		}

		result = append(result, appDef)
	}
	return result, nil
}

func (urset AppRoleDefinitionSet) IsZero() bool {
	return len(urset) == 0
}

func (urset AppRoleDefinitionSet) String() string {
	var sb strings.Builder
	for i, appDef := range urset {
		if i > 0 {
			sb.WriteString(";")
		}

		sb.WriteString(fmt.Sprintf("%s:%s:%s", appDef.AppId, appDef.AppType, appDef.AppRoles))
	}

	return sb.String()
}

func (urset AppRoleDefinitionSet) MatchTypeAndIndex(appId, appType, role string) (int, int) {
	match := MatchTypeNone
	matchNdx := -1
	for i, urd := range urset {
		mt := urd.MatchType(appId, appType, role)
		if mt > match {
			match = mt
			matchNdx = i
		}

	}
	return match, matchNdx
}

func (urset AppRoleDefinitionSet) Slice(i int) AppRoleDefinitionSet {

	if i < 0 || i >= len(urset) {
		return urset
	}

	if len(urset) == 1 {
		return nil
	}

	var sliced AppRoleDefinitionSet
	if i > 0 {
		sliced = append(sliced, urset[:i]...)
	}

	if i < len(urset)-2 {
		sliced = append(sliced, urset[i+1:]...)
	}

	return sliced
}

func (urset AppRoleDefinitionSet) MatchRole(appId, appType, role string) bool {

	mt, _ := urset.MatchTypeAndIndex(appId, appType, role)
	if mt == MatchTypeRole || (mt == MatchTypeAppType && role == "any") {
		return true
	}

	return false
}

// @tpm-schematics:end-region("bottom-file-section")
