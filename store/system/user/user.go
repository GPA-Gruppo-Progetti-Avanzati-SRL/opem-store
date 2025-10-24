package user

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
import "github.com/rs/zerolog/log"

const (
	EntityType    = "user"
	StatusActive  = "active"
	StatusBlocked = "blocked"
)

// @tpm-schematics:end-region("top-file-section")

type User struct {
	OId            bson.ObjectID         `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Nickname       string                `json:"nickname,omitempty" bson:"nickname,omitempty" yaml:"nickname,omitempty"`
	Et             string                `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Firstname      string                `json:"firstname,omitempty" bson:"firstname,omitempty" yaml:"firstname,omitempty"`
	Lastname       string                `json:"lastname,omitempty" bson:"lastname,omitempty" yaml:"lastname,omitempty"`
	Email          string                `json:"email,omitempty" bson:"email,omitempty" yaml:"email,omitempty"`
	Password       string                `json:"password,omitempty" bson:"password,omitempty" yaml:"password,omitempty"`
	Roles          []commons.UserRole    `json:"roles,omitempty" bson:"roles,omitempty" yaml:"roles,omitempty"`
	SysInfo        commons.SysInfo       `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`
	ProfilePicture commons.FileReference `json:"profile_picture,omitempty" bson:"profile_picture,omitempty" yaml:"profile_picture,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s User) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Nickname == "" && s.Et == "" && s.Firstname == "" && s.Lastname == "" && s.Email == "" && s.Password == "" && len(s.Roles) == 0 && s.SysInfo.IsZero() && s.ProfilePicture.IsZero()
}

type QueryResult struct {
	Records int    `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []User `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

func (s User) HasRole4DomainSiteAppId(domain, site, appId, appType, role string) bool {
	return AnyRole4DomainSiteAppId(s.Roles, domain, site, appId, appType, role)
}

func AnyRole4DomainSiteAppId(roles []commons.UserRole, domain, site, appId, appType, role string) bool {
	const semLogContext = "user::any-role"
	for _, r := range roles {

		if !r.MatchDomainAndSite(domain, site) {
			continue
		}

		//hr := true

		//if hr && r.Domain != "*" {
		//	if r.Domain != domain {
		//		hr = false
		//	}
		//}
		//
		//if hr && r.Site != "*" {
		//	if r.Site != site {
		//		hr = false
		//	}
		//}

		roleSet, err := r.ParseApps()
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			continue
		}

		if roleSet.MatchRole(appId, appType, role) {
			return true
		}
		//if mt ==
		//
		//if hr && r.Apps != "*" {
		//	appsArray := strings.Split(r.Apps, ";")
		//
		//	hr = false
		//	for _, a := range appsArray {
		//
		//		a = strings.TrimSpace(a)
		//		if a == "" {
		//			continue
		//		}
		//
		//		if a == "*" {
		//			hr = true
		//			break
		//		}
		//
		//		ndx := strings.Index(a, "*")
		//		if ndx >= 0 {
		//			na := a[:ndx]
		//			if strings.HasPrefix(a, na) {
		//				hr = true
		//			}
		//		} else {
		//			if a == appId {
		//				hr = true
		//			}
		//		}
		//
		//		if hr == true {
		//			break
		//		}
		//	}
		//}
		//
		//if hr {
		//	return true
		//}
	}

	return false
}

//
//func (u User) SetRole(domain, site, appId, appType, appRole string) error {
//	const semLogContext = "user::set-role"
//
//	foundRoleNdx := -1
//	for i, role := range u.Roles {
//		if role.MatchDomainAndSite(domain, site) {
//			foundRoleNdx = i
//			break
//		}
//	}
//
//	if foundRoleNdx < 0 {
//		u.Roles = append(u.Roles, commons.UserRole{
//			Domain: domain,
//			Site:   site,
//			Apps:   fmt.Sprintf("%s:%s:%s", appId, appType, appRole),
//		})
//	} else {
//		appRoleSet, err := u.Roles[foundRoleNdx].ParseApps()
//		if err != nil {
//			log.Error().Err(err).Msg(semLogContext)
//		}
//
//		if appRoleSet.IsZero() {
//			u.Roles = append(u.Roles, commons.UserRole{
//				Domain: domain,
//				Site:   site,
//				Apps:   fmt.Sprintf("%s:%s:%s", appId, appType, appRole),
//			})
//		} else {
//			mt, mndx := appRoleSet.MatchType(appId, appType, appRole)
//			switch mt {
//			case commons.MatchTypeNone:
//				appRoleSet = append(appRoleSet, commons.AppRoleDefinition{
//					AppId:    appId,
//					AppType:  appType,
//					AppRoles: appRole,
//				})
//			case commons.MatchTypeAppId:
//				appRoleSet[mndx].AppType = "*"
//				appRoleSet[mndx].AppRoles = appRole
//			case commons.MatchTypeAppType:
//				appRoleSet[mndx].AppRoles = appRole
//			case commons.MatchTypeRole:
//			default:
//				log.Error().Int("match-type", mt).Msg(semLogContext + " - unexpected app role match type")
//			}
//		}
//	}
//
//	return nil
//}

// @tpm-schematics:end-region("bottom-file-section")
