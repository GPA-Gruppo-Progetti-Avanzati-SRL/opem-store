#!/bin/sh
# export mongoServer=mongodb://localhost:27017/opem
export mongoServer=mongodb+srv://marioimperato:Y7csjKg4l4c1egiA@maidevfundamentalsclust.x5i5iii.mongodb.net/opem


mongosh $mongoServer  --file apicore-domain.js
mongosh $mongoServer  --file apicore-user.js
mongosh $mongoServer  --file apicore-kv.js
mongosh $mongoServer  --file apicore-session.js

mongosh $mongoServer --file fix-migrator-process.js
mongosh $mongoServer --file create-indexes.js
