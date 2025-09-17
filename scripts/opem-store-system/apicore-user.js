const opemDbName = "opem"
const opemCollectionName = "apicore_user"

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
