package keyvaluepackage

import (
	"context"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindByDomainSiteName(c *mongo.Collection, domain, site, pkgName string) (KeyValuePackage, bool, error) {
	f := filterByDomainSiteNameCategories(domain, site, pkgName, nil)

	kvp := KeyValuePackage{}

	crs, err := c.Find(context.TODO(), f.Build(), nil)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return kvp, false, err
		}

		return kvp, false, nil
	}

	for crs.Next(context.TODO()) {
		var kvp1 KeyValuePackage
		if err = crs.Decode(&kvp1); err != nil {
			return kvp, false, err
		}

		if ok, err := kvp1.IsMoreSpecificThan(&kvp); ok {
			kvp = kvp1
		} else if err != nil {
			return kvp, false, err
		}
	}

	if err = crs.Err(); err != nil {
		return kvp, false, err
	}

	if !kvp.IsZero() {
		scopeType, _ := ScopeTypeAndPathFromDomainSite(domain, site)
		kvpScopeType, _ := kvp.ScopeType()
		if scopeType != kvpScopeType {
			kvp.Inherited = true
		}
	}

	return kvp, !kvp.IsZero(), nil
}

func filterByDomainSiteNameCategories(domain, site, pkgName string, categories []string) Filter {
	f := Filter{}
	criteria := f.Or().AndNameEqTo(pkgName).AndCategoryIn(categories)

	if domain == store.RootDomain {
		criteria.AndScopeEqTo(store.RootDomain)
	} else {
		if site == store.SiteWildCard {
			criteria.AndScopeIn([]string{
				store.RootDomain,
				fmt.Sprintf("%s/%s", store.RootDomain, domain)})
		} else {
			criteria.AndScopeIn([]string{
				store.RootDomain,
				fmt.Sprintf("%s/%s", store.RootDomain, domain),
				fmt.Sprintf("%s/%s/%s", store.RootDomain, domain, site),
			})
		}
	}

	return f
}

func FindByDomainSiteCategoryList(c *mongo.Collection, domain, site string, categories []string) ([]KeyValuePackage, error) {
	reqScopeType, _ := ScopeTypeAndPathFromDomainSite(domain, site)

	f := filterByDomainSiteNameCategories(domain, site, "", categories)

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{NameFieldName, -1}})
	crs, err := c.Find(context.TODO(), f.Build(), findOptions)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}

		return nil, nil
	}

	var res []KeyValuePackage
	var kvp KeyValuePackage
	for crs.Next(context.TODO()) {
		var kvp1 KeyValuePackage
		if err = crs.Decode(&kvp1); err != nil {
			return res, err
		}

		if kvp1.Name != kvp.Name {
			if !kvp.IsZero() {
				kvpScopeType, _ := kvp.ScopeType()
				if reqScopeType != kvpScopeType {
					kvp.Inherited = true
				}
				res = append(res, kvp)
			}
			kvp = kvp1
		} else {
			if ok, err := kvp1.IsMoreSpecificThan(&kvp); ok {
				kvp = kvp1
			} else if err != nil {
				return res, err
			}
		}
	}

	if !kvp.IsZero() {
		kvpScopeType, _ := kvp.ScopeType()
		if reqScopeType != kvpScopeType {
			kvp.Inherited = true
		}
		res = append(res, kvp)
	}

	if err = crs.Err(); err != nil {
		return res, err
	}

	if !kvp.IsZero() {
		scopeType, _ := ScopeTypeAndPathFromDomainSite(domain, site)
		kvpScopeType, _ := kvp.ScopeType()
		if scopeType != kvpScopeType {
			kvp.Inherited = true
		}
	}

	return res, nil
}
