package model

import (
	"fmt"
	"strings"

	opemstore "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/commons"
	"gopkg.in/yaml.v3"
)

type KeyValue struct {
	Key         string `json:"key,omitempty" bson:"key,omitempty" yaml:"key,omitempty"`
	Value       any    `json:"value,omitempty" bson:"value,omitempty" yaml:"value,omitempty"`
	Order       int32  `json:"order,omitempty" bson:"order,omitempty" yaml:"order,omitempty"`
	Kind        string `json:"kind,omitempty" bson:"kind,omitempty" yaml:"kind,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Hint        string `json:"hint,omitempty" bson:"hint,omitempty" yaml:"hint,omitempty"`
	Icon        string `json:"icon,omitempty" bson:"icon,omitempty" yaml:"icon,omitempty"`
	Status      string `json:"status,omitempty" bson:"status,omitempty" yaml:"status,omitempty"`
	SysName     string `json:"sys_name,omitempty" bson:"sys_name,omitempty" yaml:"sys_name,omitempty"`
}

func (s KeyValue) IsZero() bool {
	return s.Key == "" && s.Value == nil && s.Order == 0 && s.Kind == "" && s.Name == "" && s.Description == "" && s.Hint == "" && s.Icon == "" && s.Status == "" && s.SysName == ""
}

type KeyValuePackage struct {
	OId         string          `json:"_id,omitempty" bson:"_id,omitempty" yaml:"_id,omitempty"`
	Bid         string          `json:"_bid,omitempty" bson:"_bid,omitempty" yaml:"_bid,omitempty"`
	Scope       string          `json:"scope,omitempty" bson:"scope,omitempty" yaml:"scope,omitempty"`
	Et          string          `json:"_et,omitempty" bson:"_et,omitempty" yaml:"_et,omitempty"`
	Category    string          `json:"category,omitempty" bson:"category,omitempty" yaml:"category,omitempty"`
	IsSystem    bool            `json:"is_system,omitempty" bson:"is_system,omitempty" yaml:"is_system,omitempty"`
	Description string          `json:"description,omitempty" bson:"description,omitempty" yaml:"description,omitempty"`
	Inherited   bool            `json:"inherited,omitempty" bson:"inherited,omitempty" yaml:"inherited,omitempty"`
	Properties  []KeyValue      `json:"properties,omitempty" bson:"properties,omitempty" yaml:"properties,omitempty"`
	SysInfo     commons.SysInfo `json:"sys_info,omitempty" bson:"sys_info,omitempty" yaml:"sys_info,omitempty"`
}

func (k *KeyValuePackage) UnmarshalYAML(value *yaml.Node) error {
	type rawKVP KeyValuePackage
	var raw rawKVP
	if err := value.Decode(&raw); err != nil {
		return err
	}
	*k = KeyValuePackage(raw)
	if k.OId == "" {
		k.OId = newOID()
	}
	return nil
}

func (s KeyValuePackage) IsZero() bool {
	return s.OId == "" && s.Bid == "" && s.Scope == "" && s.Et == "" && s.Category == "" && !s.IsSystem && s.Description == "" && !s.Inherited && len(s.Properties) == 0 && s.SysInfo.IsZero()
}

func (kvp *KeyValuePackage) ScopeType() (string, error) {
	return ScopeTypeFrom(kvp.Scope)
}

func (kvp *KeyValuePackage) IsMoreSpecificThan(another *KeyValuePackage) (bool, error) {
	return ScopeIsMoreSpecificThan(kvp.Scope, another.Scope)
}

type KVQueryResult struct {
	Records int               `json:"records,omitempty" bson:"records,omitempty" yaml:"records,omitempty"`
	Data    []KeyValuePackage `json:"data,omitempty" bson:"data,omitempty" yaml:"data,omitempty"`
}

func ScopeTypeFrom(scope string) (string, error) {
	if scope == "" {
		return "unknown-scope", fmt.Errorf("scope cannot be resolved since is missing")
	}
	if scope == opemstore.RootDomain {
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
	if domain == opemstore.RootDomain {
		return "root-scope", opemstore.RootDomain
	}
	if site == opemstore.SiteWildCard {
		return "domain-scope", strings.Join([]string{opemstore.RootDomain, domain}, "/")
	}
	return "site-scope", strings.Join([]string{opemstore.RootDomain, domain, site}, "/")
}

func ScopeIsMoreSpecificThan(scope string, another string) (bool, error) {
	if scope == "" {
		return false, nil
	}
	if another == "" {
		return true, nil
	}
	if strings.HasPrefix(scope, another) {
		return true, nil
	}
	var err error
	if !strings.HasPrefix(another, scope) {
		err = fmt.Errorf("incompatible paths compared: %s against %s", scope, another)
	}
	return false, err
}
