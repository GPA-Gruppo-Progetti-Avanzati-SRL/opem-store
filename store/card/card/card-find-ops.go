package card

import (
	"context"
	"errors"
	"time"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// @tpm-schematics:start-region("top-file-section")
import (
	"fmt"

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/card/person"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// @tpm-schematics:end-region("top-file-section")

// FindByPk ...
// @tpm-schematics:start-region("find-by-pk-signature-section")
func FindByPk(collection *mongo.Collection, domain, site, bidCard string, mustFind bool, findOptions *options.FindOneOptionsBuilder) (*Card, bool, error) {
	// @tpm-schematics:end-region("find-by-pk-signature-section")
	const semLogContext = "card::find-by-pk"
	// @tpm-schematics:start-region("log-event-section")
	evtTraceLog := log.Trace()
	evtErrLog := log.Error()
	// @tpm-schematics:end-region("log-event-section")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := Card{}

	f := Filter{}
	// @tpm-schematics:start-region("filter-section")
	// customize the filtering
	f.Or().AndDomainEqTo(domain).AndEtEqTo(EntityType).AndSiteEqTo(site).AndBidEqTo(bidCard)
	// @tpm-schematics:end-region("filter-section")
	fd := f.Build()
	evtTraceLog = evtTraceLog.Str("filter", util.MustToExtendedJsonString(fd, false, false))
	err := collection.FindOne(ctx, fd, findOptions).Decode(&ent)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		evtErrLog.Err(err).Msg(semLogContext)
		return &ent, false, err
	} else {
		if err != nil {
			if mustFind {
				evtTraceLog.Msg(semLogContext + " document not found")
				return &ent, false, err
			}

			evtTraceLog.Msg(semLogContext + " document not found but allowed")
			return &ent, false, nil
		} else {
			evtTraceLog.Msg(semLogContext + " document found")
		}
	}

	return &ent, true, nil
}

func Find(collection *mongo.Collection, f *Filter, withCount bool, findOptions *options.FindOptionsBuilder) (QueryResult, error) {
	const semLogContext = "card::find"
	fd := f.Build()
	evtTraceLog := log.Trace().Str("filter", util.MustToExtendedJsonString(fd, false, false))
	evtErrLog := log.Error().Str("filter", util.MustToExtendedJsonString(fd, false, false))
	evtTraceLog.Msg(semLogContext)

	qr := QueryResult{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if withCount {
		countDocsOptions := options.Count()
		nr, err := collection.CountDocuments(ctx, fd, countDocsOptions)
		if err != nil {
			evtErrLog.Err(err).Msg(semLogContext)
			return qr, err
		}

		qr.Records = int(nr)
	}

	cur, err := collection.Find(ctx, fd, findOptions)
	if err != nil {
		evtErrLog.Err(err).Msg(semLogContext)
		return qr, err
	}

	for cur.Next(context.Background()) {
		dto := Card{}
		err = cur.Decode(&dto)
		if err != nil {
			return qr, err
		}

		qr.Data = append(qr.Data, dto)
	}

	if cur.Err() != nil {
		return qr, cur.Err()
	}

	return qr, nil
}

// @tpm-schematics:start-region("bottom-file-section")

func FindOneByAggregationView(collection *mongo.Collection, collectionsCfg map[string]mongolks.CollectionCfg, dominio, site, bidCard string, mustFind bool) (*Card, bool, error) {
	const semLogContext = "card::find-one-by-aggregation-view"

	limit := int64(1)
	findOptions := options.FindOptions{
		Limit: &limit,
	}

	f := Filter{}
	f.Or().AndEtEqTo(EntityType).AndBidEqTo(bidCard).AndDomainEqTo(dominio).AndSiteEqTo(site)
	fd := f.Build()
	log.Trace().Str("filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)

	qr, err := FindByAggregationView(collection, collectionsCfg, &f, false, &findOptions)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, false, err
	}

	if len(qr.Data) == 0 {
		if mustFind {
			return nil, false, errors.New("document not found")
		}
		return nil, false, nil
	}

	b := qr.Data[0]
	return &b, true, nil
}

func FindByAggregationView(collection *mongo.Collection, collectionsCfg map[string]mongolks.CollectionCfg, f *Filter, withCount bool, findOptions *options.FindOptions) (QueryResult, error) {
	const semLogContext = "card::find-by-aggregation-view"

	personCollectionCfg := collectionsCfg[person.CollectionId]

	if personCollectionCfg.Name == "" {
		err := errors.New("cannot resolve collection name for personCollectionCfg")
		log.Error().Err(err).
			Str("person-collection-id", person.CollectionId).Str("person-collection-name", personCollectionCfg.Name).
			Msg(semLogContext)
		return QueryResult{}, err
	}

	fd := f.Build()
	evtTraceLog := log.Trace().Str("filter", util.MustToExtendedJsonString(fd, false, false))
	evtErrLog := log.Error().Str("filter", util.MustToExtendedJsonString(fd, false, false))
	evtTraceLog.Msg(semLogContext)

	qr := QueryResult{}

	// TODO
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	if withCount {
		countDocsOptions := options.Count()
		nr, err := collection.CountDocuments(ctx, fd, countDocsOptions)
		if err != nil {
			evtErrLog.Err(err).Msg(semLogContext)
			return qr, err
		}

		qr.Records = int(nr)
	}

	pipeline := mongo.Pipeline{}
	pipeline = append(pipeline, bson.D{{"$match", fd}})
	if findOptions != nil {
		if findOptions.Skip != nil {
			pipeline = append(pipeline, bson.D{{"$skip", findOptions.Skip}})
		}
		if findOptions.Limit != nil {
			pipeline = append(pipeline, bson.D{{"$limit", findOptions.Limit}})
		}
	}
	pipeline = append(pipeline, bson.D{
		{"$lookup", bson.D{
			{"from", personCollectionCfg.Name},
			{"let", bson.D{
				{"domain", "$domain"}, {"site", "$site"},
				{"et", person.EntityType},
				{"bid", bson.D{{"$ifNull", bson.A{"$holder.bid", "ND"}}}},
			},
			},
			{"pipeline", bson.A{
				bson.D{
					{"$match", bson.D{
						{"$expr", bson.D{
							{"$and",
								bson.A{
									bson.D{{"$eq", bson.A{"$domain", "$$domain"}}},
									bson.D{{"$eq", bson.A{"$site", "$$site"}}},
									bson.D{{"$eq", bson.A{"$_et", "$$et"}}},
									bson.D{{"$eq", bson.A{"$_bid", "$$bid"}}},
								},
							},
						}},
					}},
				},
			}},
			{"as", "doc_person"},
		}},
	})
	pipeline = append(pipeline, bson.D{
		{"$project", bson.D{
			{"domain", 1},
			{"site", 1},
			{"_bid", 1},
			{"_et", 1},
			{"card_number", 1},
			{"card_type", 1},
			{"corporate_code", 1},
			{"status", 1},
			{"id_card_ext", 1},
			{"envelope_number", 1},
			{"funct", 1},
			{"focal_point", 1},
			{"product", 1},
			{"box", 1},
			{"holder", 1},
			{"apps", 1},
			{"addresses", 1},
			{"expires_at", 1},
			{"issue_date", 1},
			{"issue_confirmation_date", 1},
			{"act_date", 1},
			{"sys_info", 1},
			{"events", 1},
			{"doc_person", bson.D{{"$arrayElemAt", bson.A{"$doc_person", 0}}}},
		}},
	})

	opts := options.Aggregate()
	cur, err := collection.Aggregate(ctx, pipeline, opts)
	if err != nil {
		evtErrLog.Err(err).Msg(semLogContext)
		return qr, err
	}

	for _, stage := range pipeline {
		b, err := bson.MarshalExtJSON(stage, true, true)
		if err == nil {
			fmt.Println(string(b))
		}
	}

	for cur.Next(context.Background()) {
		dto := Card{}
		err = cur.Decode(&dto)
		if err != nil {
			return qr, err
		}

		qr.Data = append(qr.Data, dto)
	}

	if cur.Err() != nil {
		return qr, cur.Err()
	}

	return qr, nil
}

// @tpm-schematics:end-region("bottom-file-section")
