### Schema session

#### Collection schema definition.

    {
	  "name": "session",
	  "properties": {
	    "folder-path": "./session",
	    "struct-name": "Session"
	  },
	  "attributes": [
	    {
	      "name": "oId",
	      "type": "object-id",
	      "tags": [
	        "json",
	        "_id",
	        "bson",
	        "_id"
	      ],
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "userid",
	      "type": "string",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "nickname",
	      "type": "string",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "remoteaddr",
	      "type": "string",
	      "queryable": true,
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "flags",
	      "type": "string",
	      "struct-ref": {
	        "Package": "",
	        "Item": null
	      }
	    },
	    {
	      "name": "sysinfo",
	      "type": "ref-struct",
	      "struct-ref": {
	        "struct-name": "SysInfo",
	        "is-external": true,
	        "Package": "github.com/mario-imperato/r3ds9-mongodb/model/r3ds9-apicore/commons",
	        "Item": null
	      }
	    }
	  ]
	}

