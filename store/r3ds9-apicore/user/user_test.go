package user_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/r3ds9-apicore/user"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const (
	targetDomain = "cvf"
)

func TestCache(t *testing.T) {

	r := user.NewCacheResolver(collection)

	user.NewCache(5*time.Second, 10*time.Minute)
	d, ok := user.GetFromCache(r, targetDomain)
	t.Log(d, ok)

	t.Log("briefly sleeping....")
	time.Sleep(2 * time.Second)
	d, ok = user.GetFromCache(r, targetDomain)
	t.Log(d, ok)

	t.Log("sleeping longer....")
	time.Sleep(10 * time.Second)

	d, ok = user.GetFromCache(r, targetDomain)
	t.Log(d, ok)
}

type roleTestCases struct {
	domain   string
	site     string
	appId    string
	appType  string
	appRole  string
	roles    commons.UserRoleList
	expected bool
}

func TestHasRole(t *testing.T) {

	cases := []roleTestCases{
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{}, false},
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{{Domain: "*", Site: "*", Apps: "appero;app-2"}}, false},
		{"dom1", "site1", "app-1", "admin", "role-required", commons.UserRoleList{{Domain: "*", Site: "*", Apps: "appero;app-1:admin:role-required"}}, true},
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{{Domain: "*", Site: "*", Apps: "appero;*"}}, true},
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{{Domain: "*", Site: "site2", Apps: "*"}}, false},
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{{Domain: "dom2", Site: "*", Apps: "*"}}, false},
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{{Domain: "dom1", Site: "*", Apps: "*"}}, true},
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{{Domain: "dom1", Site: "site1", Apps: "*"}, {Domain: "dom2", Site: "site2", Apps: "app-2"}}, true},
	}

	for i, c := range cases {
		ok := user.AnyRole4DomainSiteAppId(c.roles, c.domain, c.site, c.appId, c.appType, c.appRole)
		require.Equal(t, c.expected, ok, "case %d", i)
	}
}

func TestSetRole(t *testing.T) {

	cases := []roleTestCases{
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{}, false},
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{{Domain: "*", Site: "*", Apps: "appero;app-2"}}, false},
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{{Domain: "*", Site: "*", Apps: "appero;app-1:ui:role-required"}}, false},
		{"dom1", "site1", "app-1", "admin", "admin", commons.UserRoleList{{Domain: "*", Site: "*", Apps: "appero;app-1:admin:admin"}}, false},
		//{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "*", Site: "*", Apps: "appero, app-1"}}, true},
		//{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "*", Site: "*", Apps: "appero,app-*"}}, true},
		//{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "*", Site: "site2", Apps: "app-*"}}, false},
		//{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "dom2", Site: "*", Apps: "app-*"}}, false},
		//{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "dom1", Site: "*", Apps: "app-*"}}, true},
		//{"dom1", "site1", "app-1", []commons.UserRole{{Domain: "dom1", Site: "site1", Apps: "app-*"}, {Domain: "dom2", Site: "site2", Apps: "app-2"}}, true},
	}

	for _, c := range cases {
		lst, err := c.roles.Add(c.domain, c.site, c.appId, "admin", "role-required")
		require.NoError(t, err)
		t.Log(lst)
	}
}
