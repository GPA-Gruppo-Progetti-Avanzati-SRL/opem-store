const opemDbName = "opem"

let conn = db.getMongo();
let db = conn.getDB(opemDbName);

print("[system] session collection - #docs: ", db.apicore_session.countDocuments())
print("[system] domain collection - #docs: ", db.apicore_domain.countDocuments())
print("[system] site collection - #docs: ", db.apicore_site.countDocuments())
print("[system] user collection - #docs: ", db.apicore_user.countDocuments())
print("[system] kv collection - #docs: ", db.apicore_kv.countDocuments())

db.apicore_user.createIndex({ "firstname": "text", "lastname": "text", "nickname": "text" })
