const r3ds9DbName = "opem"

let conn = db.getMongo();
let db = conn.getDB(r3ds9DbName);

print("[apicms] file collection - #docs: ", db.apicms_file.countDocuments())

