/*
 * This is a testing file and populates a few documents for checking the key-value-package inheritance mechanism.
 */

const r3ds9DbName = "r3ds9"
const r3ds9CollectionName = "apicore_kv"

let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

let c = db[r3ds9CollectionName]
if (!c)  {
    db.createCollection(r3ds9CollectionName)
}
else
{
    db.apicore_kv.deleteMany({ "name": { $in : ["test-kv1", "test-kv2", "test-kv3"] }  })
}

c.insertOne(
    {
        "name": "test-kv1",
        "scope": "root",
        "objType": "kv",
        "category": "test-kv-cat1",
        "description": "Key value package (1) root level"
        ,"properties": [
            { "key": "theme", "value": "theme-1" },
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

c.insertOne(
    {
        "name": "test-kv2",
        "scope": "root",
        "objType": "kv",
        "category": "test-kv-cat1",
        "description": "Key value package (2) root level"
        ,"properties": [
            { "key": "theme", "value": "theme-1" },
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

c.insertOne(
    {
        "name": "test-kv3",
        "scope": "root",
        "objType": "kv",
        "category": "test-kv-cat1",
        "description": "Key value package (3) root level"
        ,"properties": [
            { "key": "theme", "value": "theme-1" },
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

c.insertOne(
    {
        "name": "test-kv4",
        "scope": "root",
        "objType": "kv",
        "category": "test-kv-cat2",
        "description": "Key value package (4) root level"
        ,"properties": [
            { "key": "theme", "value": "theme-1" },
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

// Override at the domain level
c.insertOne(
    {
        "name": "test-kv2",
        "scope": "root/cvf",
        "objType": "kv",
        "category": "test-kv-cat1",
        "description": "Key value package (2) overridden domain level"
        ,"properties": [
            { "key": "theme", "value": "theme-1" },
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

c.insertOne(
    {
        "name": "test-kv3",
        "scope": "root/cvf",
        "objType": "kv",
        "category": "test-kv-cat1",
        "description": "Key value package (3) overridden domain level"
        ,"properties": [
            { "key": "theme", "value": "theme-1" },
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

// Override at the site level
c.insertOne(
    {
        "name": "test-kv3",
        "scope": "root/cvf/champ42",
        "objType": "kv",
        "category": "test-kv-cat1",
        "description": "Key value package (3) overridden site level"
        ,"properties": [
            { "key": "theme", "value": "theme-1" },
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

