const opemDbName = "opem"

let conn = db.getMongo();
let db = conn.getDB(opemDbName);

print("[cms] file collection - #docs: ", db.apicms_file.countDocuments())

