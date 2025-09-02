const r3ds9DbName = "opem"
const r3ds9CollectionName = "apicore_session"

let conn = db.getMongo();
let db = conn.getDB(r3ds9DbName);

let c = db[r3ds9CollectionName]
if (!c)  {
    db.createCollection(r3ds9CollectionName)
}
else
{
    c.deleteMany({});
}
