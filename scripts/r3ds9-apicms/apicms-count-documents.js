const r3ds9DbName = "r3ds9"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

print("[apicms] file collection - #docs: ", db.apicms_file.countDocuments())

