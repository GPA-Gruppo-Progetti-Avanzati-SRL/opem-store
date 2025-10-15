let crs = db["opem_prodotto_funct"].find();
while ( crs.hasNext() ) {
    doc = crs.next()
    db.opem_prodotto.updateOne({ "_et": "PRODOTTO", "_bid": doc.prodotto._bid}, { "$push": { "params": {"params.func" : doc.func, "params.value": doc.valParam, "params.mysql.prog_param": doc.progParam }} })
}

db.opem_prodotto_prod_param.drop();