const opemDbName = "opem"

let conn = db.getMongo();
let db = conn.getDB(opemDbName);

print("[territorio] territorio collection - #docs: ", db.trt_territorio.countDocuments())

