const r3ds9DbName = "r3ds9"
const r3ds9CollectionName = "apicms_file"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

let c = db[r3ds9CollectionName]
if (!c)  {
    db.createCollection(r3ds9CollectionName)
}
else
{
    c.deleteMany({});
}

// let cCore = db["apicore_user"]
// let crs = cCore.find()
// while (crs.hasNext()) {
//     let u = crs.next()
//     c.insertOne(
//         {
//             "userId" : u._id.toString(),
//             "objType"  : "user-profile"
//             ,"sysinfo": {
//                 "createdat": new Date(),
//                 "modifiedat": new Date()
//             }
//         });
// }

