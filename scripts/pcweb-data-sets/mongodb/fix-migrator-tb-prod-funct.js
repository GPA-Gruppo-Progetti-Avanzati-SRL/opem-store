let crs = db["opem_prodotto_funct"].find();
while ( crs.hasNext() ) {
    doc = crs.next()
    db.opem_prodotto.updateOne(
        { "_et": "PRODOTTO", "_bid": doc.product.bid, "site": doc.site, "domain": doc.domain },
        { "$push": { "apps": doc.app } }
    )
}

db.opem_prodotto_funct.drop();