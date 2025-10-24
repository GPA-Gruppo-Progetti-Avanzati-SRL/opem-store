print(new Date() + " #################### - fix-mig-05-opem-user.js")

const opemCollectionName = "opem_user"
db[opemCollectionName].insertOne(
    {
        "nickname" : "root",
        "_et"  : "root-user",
        "password" : "8a3308114f80796bb4b6d407e605752b167c6b75",
        "firstname": "G",
        "lastname" : "Root",
        "roles": [
            {
                "domain": "*"
                ,"site" : "*"
                ,"apps" : "*"
            }
        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    });

db[opemCollectionName].insertOne(
    {
        "nickname" : "guest",
        "_et"  : "guest-user",
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    });

db[opemCollectionName].insertOne(
    {
        "nickname" : "nambalorian",
        "_et"  : "user",
        "password" : "8a3308114f80796bb4b6d407e605752b167c6b75",
        "firstname": "Wanamba",
        "lastname" : "ForEver",
        "roles": [
            {
                "domain": "*"
                ,"site" : "*"
                ,"apps" : "app-home:admin:requested-role;app-people-admin:admin"
            }
        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    });
