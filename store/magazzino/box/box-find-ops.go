package box

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

	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/magazzino/magazzino"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/opem-store/store/prodotto/prodotto"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// @tpm-schematics:end-region("top-file-section")

// FindByPk ...
// @tpm-schematics:start-region("find-by-pk-signature-section")
func FindByPk(collection *mongo.Collection, dominio, site, bidMagazzino, bidBox string, mustFind bool, findOptions *options.FindOneOptionsBuilder) (*Box, bool, error) {
	// @tpm-schematics:end-region("find-by-pk-signature-section")
	const semLogContext = "box::find-by-pk"
	// @tpm-schematics:start-region("log-event-section")
	evtTraceLog := log.Trace()
	evtErrLog := log.Error()
	// @tpm-schematics:end-region("log-event-section")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ent := Box{}

	f := Filter{}
	// @tpm-schematics:start-region("filter-section")
	// customize the filtering
	f.Or().AndEtEqTo(EntityType).AndBidEqTo(bidBox).AndBidMagazzinoEqTo(bidMagazzino).AndDomainEqTo(dominio).AndSiteEqTo(site)
	// @tpm-schematics:end-region("filter-section")
	fd := f.Build()
	evtTraceLog = evtTraceLog.Str("filter", util.MustToExtendedJsonString(fd, false, false))
	err := collection.FindOne(ctx, fd, findOptions).Decode(&ent)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		evtErrLog.Err(err).Msg(semLogContext)
		return nil, false, err
	} else {
		if err != nil {
			if mustFind {
				evtTraceLog.Msg(semLogContext + " document not found")
				return nil, false, err
			}

			evtTraceLog.Msg(semLogContext + " document not found but allowed")
			return nil, false, nil
		} else {
			evtTraceLog.Msg(semLogContext + " document found")
		}
	}

	return &ent, true, nil
}

func FindFirst(collection *mongo.Collection, f *Filter, findOptions *options.FindOptionsBuilder) (*Box, error) {
	const semLogContext = "box::find-first"
	fd := f.Build()
	log.Trace().Str("filter", util.MustToExtendedJsonString(fd, false, false)).Msg(semLogContext)

	cur, err := collection.Find(context.Background(), fd, findOptions)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	if cur.Next(context.Background()) {
		dto := Box{}
		err = cur.Decode(&dto)
		if err != nil {
			return nil, err
		}

		return &dto, nil
	}

	return nil, nil
}

func Find(collection *mongo.Collection, f *Filter, withCount bool, findOptions *options.FindOptionsBuilder) (QueryResult, error) {
	const semLogContext = "box::find"
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
		dto := Box{}
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

func FindOneByAggregationView(collection *mongo.Collection, collectionsCfg map[string]mongolks.CollectionCfg, dominio string, site string, bidMagazzino string, bidBox string, mustFind bool) (*Box, bool, error) {
	const semLogContext = "box::find-one-by-aggregation-view"

	limit := int64(1)
	findOptions := options.FindOptions{
		Limit: &limit,
	}

	f := Filter{}
	f.Or().AndEtEqTo(EntityType).AndBidEqTo(bidBox).AndBidMagazzinoEqTo(bidMagazzino).AndDomainEqTo(dominio).AndSiteEqTo(site)
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
	const semLogContext = "box::find-by-aggregation-view"

	magazzinoCollectionCfg := collectionsCfg[magazzino.CollectionId]
	prodottoCollectionCfg := collectionsCfg[prodotto.CollectionId]
	if magazzinoCollectionCfg.Name == "" || prodottoCollectionCfg.Name == "" {
		err := errors.New("cannot resolve collections name for magazzino or prodotto")
		log.Error().Err(err).
			Str("magazzino-collection-id", magazzino.CollectionId).Str("magazzino-collection-name", magazzinoCollectionCfg.Name).
			Str("prodotto-collection-id", prodotto.CollectionId).Str("prodotto-collection-name", prodottoCollectionCfg.Name).
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
			{"from", magazzinoCollectionCfg.Name},
			{"let", bson.D{
				{"domain", "$domain"}, {"site", "$site"},
				{"et", magazzino.EntityType}, {"bid", "$magazzino.bid"},
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
			{"as", "doc_magazzino"},
		}},
	})
	pipeline = append(pipeline, bson.D{
		{"$lookup", bson.D{
			{"from", prodottoCollectionCfg.Name},
			{"let", bson.D{
				{"domain", "$domain"}, {"site", "$site"},
				{"et", prodotto.EntityType}, {"bid", "$prodotto.bid"},
			},
			},
			{"pipeline", bson.A{
				bson.D{
					{"$match", bson.D{
						{"$expr", bson.D{
							{"$and", bson.A{
								bson.D{{"$eq", bson.A{"$domain", "$$domain"}}},
								bson.D{{"$eq", bson.A{"$site", "$$site"}}},
								bson.D{{"$eq", bson.A{"$_et", "$$et"}}},
								bson.D{{"$eq", bson.A{"$_bid", "$$bid"}}},
							}},
						}},
					}},
				},
			}},
			{"as", "doc_prodotto"},
		}},
	})
	pipeline = append(pipeline, bson.D{
		{"$project", bson.D{
			{"_bid", 1},
			{"_et", 1},
			{"domain", 1},
			{"site", 1},
			{"info", 1},
			{"status", 1},
			{"recipient", 1},
			{"events", 1},
			{"activities", 1},
			{"notes", 1},
			{"sys_info", 1},
			{"card_bids_range", 1},
			{"doc_magazzino", bson.D{{"$arrayElemAt", bson.A{"$doc_magazzino", 0}}}},
			{"doc_prodotto", bson.D{{"$arrayElemAt", bson.A{"$doc_prodotto", 0}}}},
		}},
	})
	pipeline = append(pipeline, bson.D{
		{"$project", bson.D{
			{"_bid", 1},
			{"_et", 1},
			{"domain", 1},
			{"site", 1},
			{"info", 1},
			{"status", 1},
			{"recipient", 1},
			{"events", 1},
			{"activities", 1},
			{"notes", 1},
			{"supply_type", 1},
			{"sys_info", 1},
			{"card_bids_range", 1},
			{"magazzino.bid", "$doc_magazzino._bid"},
			{"prodotto.bid", "$doc_prodotto._bid"},
			{"prodotto.text", "$doc_prodotto.name"},
			// {"prodotto.text", bson.D{{"$concat", bson.A{"$doc_prodotto._bid", " - ", "$doc_prodotto.name"}}}},
			{"focal_point.bid", "$doc_magazzino.focal_point.bid"},
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
		dto := Box{}
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
