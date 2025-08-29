const r3ds9DbName = "r3ds9"
const r3ds9CollectionName = "apicore_user"

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

c.insertOne(
    {
        "nickname" : "root",
        "objType"  : "root-user",
        "password" : "8a3308114f80796bb4b6d407e605752b167c6b75",
        "firstname": "G",
        "lastname" : "Root",
        "roles": [
            {
                "domain": "*"
                ,"site" : "*"
                ,"apps" : "*"
            }
        ]
        ,"sysInfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

c.insertOne(
    {
        "nickname" : "guest",
        "objType"  : "guest-user"
        ,"sysInfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

c.insertOne(
    {
        "nickname" : "nambalorian",
        "objType"  : "user",
        "password" : "8a3308114f80796bb4b6d407e605752b167c6b75",
        "firstname": "Wanamba",
        "lastname" : "ForEver",
        "roles": [
            {
                "domain": "*"
                ,"site" : "*"
                ,"apps" : "app-home:admin:requested-role;app-people-admin:admin"
            }
        ]
        ,"sysInfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });
