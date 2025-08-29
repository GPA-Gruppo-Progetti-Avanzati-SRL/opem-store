const r3ds9DbName = "r3ds9"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

print("[apicore] session collection - #docs: ", db.apicore_session.countDocuments())
print("[apicore] domain collection - #docs: ", db.apicore_domain.countDocuments())
print("[apicore] site collection - #docs: ", db.apicore_site.countDocuments())
print("[apicore] user collection - #docs: ", db.apicore_user.countDocuments())
print("[apicore] kv collection - #docs: ", db.apicore_kv.countDocuments())

db.apicore_user.createIndex({ "firstname": "text", "lastname": "text", "nickname": "text" })
