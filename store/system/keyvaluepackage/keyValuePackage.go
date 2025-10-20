package keyvaluepackage

import "go.mongodb.org/mongo-driver/v2/bson"
import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"

// @tpm-schematics:start-region("top-file-section")
import (
	"fmt"
	"strings"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store"
)

const (
	EntityType   = "kv"
	CollectionId = "key-value-package"
)

// @tpm-schematics:end-region("top-file-section")

type KeyValuePackage struct {
	OId         bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Bid         string          `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Scope       string          `json:"scope,omitempty" bson:"scope,omitempty" yaml:"scope,omitempty"`
	Et          string          `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Category    string          `json:"category,omitempty" bson:"category,omitempty" yaml:"category,omitempty"`
	IsSystem    bool            `json:"is_system,omitempty" bson:"is_system,omitempty" yaml:"is_system,omitempty"`
	Description string          `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Inherited   bool            `json:"inherited,omitempty" bson:"inherited,omitempty" yaml:"inherited,omitempty"`
	Properties  []KeyValue      `json:"properties,omitempty" bson:"properties,omitempty" yaml:"properties,omitempty"`
	SysInfo     commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`

	// @tpm-schematics:start-region("struct-section")
	// @tpm-schematics:end-region("struct-section")
}

func (s KeyValuePackage) IsZero() bool {
	return s.OId == bson.NilObjectID && s.Bid == "" && s.Scope == "" && s.Et == "" && s.Category == "" && !s.IsSystem && s.Description == "" && !s.Inherited && len(s.Properties) == 0 && s.SysInfo.IsZero()
}

type QueryResult struct {
	Records int               `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []KeyValuePackage `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

// @tpm-schematics:start-region("bottom-file-section")

func (kvp *KeyValuePackage) ScopeType() (string, error) {
	return ScopeTypeFrom(kvp.Scope)
}

func (kvp *KeyValuePackage) IsMoreSpecificThan(another *KeyValuePackage) (bool, error) {
	return ScopeIsMoreSpecificThan(kvp.Scope, another.Scope)
}

func ScopeTypeFrom(scope string) (string, error) {
	if scope == "" {
		return "unknown-scope", fmt.Errorf("scope cannot be resolved since is missing")
	}

	if scope == store.RootDomain {
		return "root-scope", nil
	}

	var scopeType string
	var err error
	switch strings.Count(scope, "/") {
	case 1:
		scopeType = "domain-scope"
	case 2:
		scopeType = "site-scope"
	default:
		err = fmt.Errorf("malformed scope")
	}

	return scopeType, err
}

func ScopeTypeAndPathFromDomainSite(domain, site string) (string, string) {
	if domain == store.RootDomain {
		return "root-scope", store.RootDomain
	}

	if site == store.SiteWildCard {
		return "domain-scope", strings.Join([]string{store.RootDomain, domain}, "/")
	}

	return "site-scope", strings.Join([]string{store.RootDomain, domain, site}, "/")
}

func ScopeIsMoreSpecificThan(scope string, another string) (bool, error) {
	if scope == "" {
		return false, nil
	}

	if another == "" {
		return true, nil
	}

	/*
		_, scopePath, err := ScopeTypeFrom(scope)
		if err != nil {
			return false, err
		}

		_, anotherScopePath, err := ScopeTypeFrom(another)
		if err != nil {
			return false, err
		}
	*/

	if strings.HasPrefix(scope, another) {
		return true, nil
	}

	var err error
	if !strings.HasPrefix(another, scope) {
		err = fmt.Errorf("incompatible paths compared: %s against %s", scope, another)
	}

	return false, err
}

// @tpm-schematics:end-region("bottom-file-section")
