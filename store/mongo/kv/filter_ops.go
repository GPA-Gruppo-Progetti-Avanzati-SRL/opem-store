package kv

import (
	"context"
	"fmt"

	opemstore "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store"
	model "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func FindByDomainSiteName(c *mongo.Collection, domain, site, pkgName string) (model.KeyValuePackage, bool, error) {
	f := filterByDomainSiteNameCategories(domain, site, pkgName, nil)

	kvp := model.KeyValuePackage{}

	crs, err := c.Find(context.TODO(), f.Build(), nil)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return kvp, false, err
		}
		return kvp, false, nil
	}

	for crs.Next(context.TODO()) {
		var kvp1 model.KeyValuePackage
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
		scopeType, _ := model.ScopeTypeAndPathFromDomainSite(domain, site)
		kvpScopeType, _ := kvp.ScopeType()
		if scopeType != kvpScopeType {
			kvp.Inherited = true
		}
	}

	return kvp, !kvp.IsZero(), nil
}

func filterByDomainSiteNameCategories(domain, site, pkgName string, categories []string) Filter {
	f := Filter{}
	criteria := f.Or().AndBidEqTo(pkgName).AndEtEqTo(EntityType).AndCategoryIn(categories)

	if domain == opemstore.RootDomain {
		criteria.AndScopeEqTo(opemstore.RootDomain)
	} else {
		if site == opemstore.SiteWildCard {
			criteria.AndScopeIn([]string{
				opemstore.RootDomain,
				fmt.Sprintf("%s/%s", opemstore.RootDomain, domain),
			})
		} else {
			criteria.AndScopeIn([]string{
				opemstore.RootDomain,
				fmt.Sprintf("%s/%s", opemstore.RootDomain, domain),
				fmt.Sprintf("%s/%s/%s", opemstore.RootDomain, domain, site),
			})
		}
	}

	return f
}

func FindByDomainSiteCategoryList(c *mongo.Collection, domain, site string, categories []string) ([]model.KeyValuePackage, error) {
	reqScopeType, _ := model.ScopeTypeAndPathFromDomainSite(domain, site)

	f := filterByDomainSiteNameCategories(domain, site, "", categories)

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: BidFieldName, Value: -1}})
	crs, err := c.Find(context.TODO(), f.Build(), findOptions)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, nil
	}

	var res []model.KeyValuePackage
	var kvp model.KeyValuePackage
	for crs.Next(context.TODO()) {
		var kvp1 model.KeyValuePackage
		if err = crs.Decode(&kvp1); err != nil {
			return res, err
		}

		if kvp1.Bid != kvp.Bid {
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

	return res, nil
}
