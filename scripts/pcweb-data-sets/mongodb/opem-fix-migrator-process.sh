#!/bin/sh
# export mongoServer=mongodb://localhost:27017/opem
export mongoServer=mongodb+srv://marioimperato:Y7csjKg4l4c1egiA@maidevfundamentalsclust.x5i5iii.mongodb.net/opem

mongosh $mongoServer --file fix-migrator-drop-rel-mig-indexes.js
mongosh $mongoServer --file create-indexes.js

mongosh $mongoServer  --file opem-system-domain.js
mongosh $mongoServer  --file apicore-user.js
mongosh $mongoServer  --file opem-system-kv.js
mongosh $mongoServer  --file apicore-session.js

mongosh $mongoServer --file fix-migrator-companies.js
mongosh $mongoServer --file fix-migrator-focal-point.js
mongosh $mongoServer --file fix-migrator-tb-dom-param.js
mongosh $mongoServer --file fix-migrator-summary.js



