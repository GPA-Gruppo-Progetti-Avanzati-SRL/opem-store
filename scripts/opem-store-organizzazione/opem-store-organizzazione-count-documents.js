const opemDbName = "opem"

let conn = db.getMongo();
let db = conn.getDB(opemDbName);

print("[organizzazione] organizzazione collection - #docs: ", db.org_organizzazione.countDocuments())

