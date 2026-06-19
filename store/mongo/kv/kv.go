package kv

import (
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
)

const (
	EntityType   = "kv"
	CollectionId = "key-value-package"
)

func ScopeTypeFrom(scope string) (string, error) {
	return model.ScopeTypeFrom(scope)
}

func ScopeTypeAndPathFromDomainSite(domain, site string) (string, string) {
	return model.ScopeTypeAndPathFromDomainSite(domain, site)
}

func ScopeIsMoreSpecificThan(scope string, another string) (bool, error) {
	return model.ScopeIsMoreSpecificThan(scope, another)
}
