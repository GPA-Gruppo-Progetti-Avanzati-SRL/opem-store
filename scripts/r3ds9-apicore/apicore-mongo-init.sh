#!/bin/bash
mongosh mongodb://localhost:27017/r3ds9  --file apicore-session.js
mongosh mongodb://localhost:27017/r3ds9  --file apicore-site.js
mongosh mongodb://localhost:27017/r3ds9  --file apicore-domain.js
mongosh mongodb://localhost:27017/r3ds9  --file apicore-user.js
mongosh mongodb://localhost:27017/r3ds9  --file apicore-session.js
mongosh mongodb://localhost:27017/r3ds9  --file apicore-kv.js
mongosh mongodb://localhost:27017/r3ds9  --file apicore-count-documents.js