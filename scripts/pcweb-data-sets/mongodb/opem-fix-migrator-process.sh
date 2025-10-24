#!/bin/sh
# export mongoServer=mongodb://localhost:27017/opem
export mongoServer=mongodb+srv://marioimperato:Y7csjKg4l4c1egiA@maidevfundamentalsclust.x5i5iii.mongodb.net/opem
# export mongoServer=mongodb+srv://gpa:gpa@mongo.gpagroup.it/opem_mig_wip?tls=false

mongosh $mongoServer --file fix-mig-01-drop-rel-mig-indexes.js
mongosh $mongoServer --file fix-mig-02-create-indexes.js
mongosh $mongoServer  --file fix-mig-03-sequences.js
mongosh $mongoServer  --file fix-mig-04-domain.js
mongosh $mongoServer  --file fix-mig-05-opem-user.js
mongosh $mongoServer  --file fix-mig-06-opem-system-kv.js
mongosh $mongoServer  --file fix-mig-07-opem-session.js

mongosh $mongoServer --file fix-mig-08-companies.js
mongosh $mongoServer --file fix-mig-09-focal-point.js
mongosh $mongoServer --file fix-mig-10-tb-dom-param.js
mongosh $mongoServer --file fix-mig-11-tb-prod-funct.js

mongosh $mongoServer --file fix-mig-12-addresses.js
mongosh $mongoServer --file fix-mig-13-cards.js
mongosh $mongoServer --file fix-mig-14-persons.js
mongosh $mongoServer --file fix-mig-99-summary.js



