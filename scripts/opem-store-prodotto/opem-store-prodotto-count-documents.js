const opemDbName = "opem"

let conn = db.getMongo();
let db = conn.getDB(opemDbName);

print("[prodotto] prodotto collection - #docs: ", db.prd_prodotto.countDocuments())

