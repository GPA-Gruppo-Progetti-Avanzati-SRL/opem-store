# Agregation box_view

## pipeline

```json
[
  {
    "$match": {
      "_et": "BOX"
    }
  },
  {
    "$skip": 20
  },
  {
    "$limit": 10
  },
  {
    "$lookup": {
      "from": "mag_magazzini",
      "let": {
        "et": "MAGAZZINO",
        "bid": "$magazzino.bid"
      },
      "pipeline": [
        {
          "$match": {
            "$expr": {
              "$and": [
                {
                  "$eq": ["$_et", "$$et"]
                },
                {
                  "$eq": ["$_bid", "$$bid"]
                }
              ]
            }
          }
        }
      ],
      "as": "doc_magazzino"
    }
  },
  {
    "$lookup": {
      "from": "prd_prodotto",
      "let": {
        "et": "PRODOTTO",
        "bid": "$prodotto.bid"
      },
      "pipeline": [
        {
          "$match": {
            "$expr": {
              "$and": [
                {
                  "$eq": ["$_et", "$$et"]
                },
                {
                  "$eq": ["$_bid", "$$bid"]
                }
              ]
            }
          }
        }
      ],
      "as": "doc_prodotto"
    }
  },
  {
    "$project": {
      "_bid": 1,
      "_et": 1,
      "domain": 1,
      "site": 1,
      "info": 1,
      "status": 1,
      "recipient": 1,
      "sys_info": 1,
      "doc_magazzino": {
        "$arrayElemAt": ["$doc_magazzino", 0]
      },
      "doc_prodotto": {
        "$arrayElemAt": ["$doc_prodotto", 0]
      }
    }
  },
  {
    "$project": {
      "_bid": 1,
      "_et": 1,
      "domain": 1,
      "site": 1,
      "info": 1,
      "status": 1,
      "recipient": 1,
      "sys_info": 1,
      "magazzino.bid": "$doc_magazzino._bid",
      "prodotto.bid": "$doc_prodotto._bid",
      "prodotto.text": {
        "$concat": [
          "$doc_prodotto._bid",
          " - ",
          "$doc_prodotto.name"
        ]
      },
      "focal_point.bid":
      "$doc_magazzino.focal_point.bid"
    }
  }
]
```

## golang conversion

```golang
_ = mongo.Pipeline{
		bson.D{{"$match", bson.D{{"_et", "BOX"}}}},
		bson.D{{"$skip", 20}},
		bson.D{{"$limit", 10}},
		bson.D{
			{"$lookup", bson.D{
				{"from", "mag_magazzini"},
				{"let", bson.D{{"et", "MAGAZZINO"}, {"bid", "$magazzino.bid"}}},
				{"pipeline", bson.A{
					bson.D{
						{"$match", bson.D{
							{"$expr", bson.D{
								{"$and",
									bson.A{
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
		},
		bson.D{
			{"$lookup", bson.D{
				{"from", "prd_prodotto"},
				{"let", bson.D{{"et", "PRODOTTO"}, {"bid", "$prodotto.bid"}}},
				{"pipeline", bson.A{
					bson.D{
						{"$match", bson.D{
							{"$expr", bson.D{
								{"$and", bson.A{
									bson.D{{"$eq", bson.A{"$_et", "$$et"}}},
									bson.D{{"$eq", bson.A{"$_bid", "$$bid"}}},
								}},
							}},
						}},
					},
				}},
				{"as", "doc_prodotto"},
			}},
		},
		bson.D{
			{"$project", bson.D{
				{"_bid", 1},
				{"_et", 1},
				{"domain", 1},
				{"site", 1},
				{"info", 1},
				{"status", 1},
				{"recipient", 1},
				{"sys_info", 1},
				{"doc_magazzino", bson.D{{"$arrayElemAt", bson.A{"$doc_magazzino", 0}}}},
				{"doc_prodotto", bson.D{{"$arrayElemAt", bson.A{"$doc_prodotto", 0}}}},
			}},
		},
		bson.D{
			{"$project", bson.D{
				{"_bid", 1},
				{"_et", 1},
				{"domain", 1},
				{"site", 1},
				{"info", 1},
				{"status", 1},
				{"recipient", 1},
				{"sys_info", 1},
				{"magazzino.bid", "$doc_magazzino._bid"},
				{"prodotto.bid", "$doc_prodotto._bid"},
				{"prodotto.text", bson.D{{"$concat", bson.A{"$doc_prodotto._bid", " - ", "$doc_prodotto.name"}}}},
				{"focal_point.bid", "$doc_magazzino.focal_point.bid"},
			}},
		},
	}
```