let kv_cursor = db["opem_system_dom_param"].aggregate(
    [
        {
            $group:
            /**
             * _id: The id of the group.
             * fieldN: The first field name.
             */
                {
                    _id: "$name",
                    properties: {
                        $push: {
                            key: "$parmKey",
                            value: "$value",
                            order: 0,
                            kind: "string"
                        }
                    }
                }
        },
        {
            $addFields:
            /**
             * newField: The new field name.
             * expression: The new field expression.
             */
                {
                    _bid: "$_id",
                    _et: "KV",
                    scope: "root/card",
                    inherited: false,
                    description: "$_id",
                    category: "lookup",
                    obj_type: "kv",
                    "sys_info.created_at": "$$NOW",
                    "sys_info.modified_at": "$$NOW",
                    "sys_info.status": "active"
                }
        },
        {
            $project:
            /**
             * specifications: The fields to
             *   include or exclude.
             */
                {
                    _id: 0
                }
        }
    ]
);

while ( kv_cursor.hasNext() ) {
    doc = kv_cursor.next()
    db.opem_system.insertOne(doc)
}

db.opem_system_dom_param.drop();