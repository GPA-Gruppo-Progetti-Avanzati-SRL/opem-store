print(new Date() + " #################### - fix-mig-03-sequences.js")

// opem_system_seq_address
let crs = db["opem_system_seq_address"].find();
while ( crs.hasNext() ) {
    doc = crs.next()
    db.opem_system.updateOne(
        { "_et": doc._et, "_bid": doc._bid, "site": doc.site, "domain": doc.domain },
        { "$set":
          {
            "_et": doc._et,
            "_bid": doc._bid,
            "site": doc.site,
            "domain": doc.domain,
            "format": doc.format,
            "value": doc.value
          }
        },
        { upsert: true}
    )
}

// opem_system_seq_appnum
crs = db["opem_system_seq_appnum"].find();
while ( crs.hasNext() ) {
    doc = crs.next()
    db.opem_system.updateOne(
        { "_et": doc._et, "_bid": doc._bid, "site": doc.site, "domain": doc.domain },
        { "$set":
                {
                    "_et": doc._et,
                    "_bid": doc._bid,
                    "site": doc.site,
                    "domain": doc.domain,
                    "format": doc.format,
                    "value": doc.value
                }
        },
        { upsert: true}
    )
}

// opem_system_seq_card
crs = db["opem_system_seq_card"].find();
while ( crs.hasNext() ) {
    doc = crs.next()
    db.opem_system.updateOne(
        { "_et": doc._et, "_bid": doc._bid, "site": doc.site, "domain": doc.domain },
        { "$set":
                {
                    "_et": doc._et,
                    "_bid": doc._bid,
                    "site": doc.site,
                    "domain": doc.domain,
                    "format": doc.format,
                    "value": doc.value
                }
        },
        { upsert: true}
    )
}

// opem_system_seq_mag
crs = db["opem_system_seq_mag"].find();
while ( crs.hasNext() ) {
    doc = crs.next()
    db.opem_system.updateOne(
        { "_et": doc._et, "_bid": doc._bid, "site": doc.site, "domain": doc.domain },
        { "$set":
                {
                    "_et": doc._et,
                    "_bid": doc._bid,
                    "site": doc.site,
                    "domain": doc.domain,
                    "format": doc.format,
                    "value": doc.value
                }
        },
        { upsert: true}
    )
}

// opem_system_seq_pers
crs = db["opem_system_seq_pers"].find();
while ( crs.hasNext() ) {
    doc = crs.next()
    db.opem_system.updateOne(
        { "_et": doc._et, "_bid": doc._bid, "site": doc.site, "domain": doc.domain },
        { "$set":
                {
                    "_et": doc._et,
                    "_bid": doc._bid,
                    "site": doc.site,
                    "domain": doc.domain,
                    "format": doc.format,
                    "value": doc.value
                }
        },
        { upsert: true}
    )
}

db.opem_system_seq_address.drop();
db.opem_system_seq_appnum.drop();
db.opem_system_seq_card.drop();
db.opem_system_seq_mag.drop();
db.opem_system_seq_pers.drop();