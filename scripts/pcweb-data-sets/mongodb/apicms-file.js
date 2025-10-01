const opemDbName = "opem"
const opemCollectionName = "apicms_file"

let conn = db.getMongo();
let db = conn.getDB(opemDbName);

let c = db[opemCollectionName]
if (!c)  {
    db.createCollection(opemCollectionName)
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

