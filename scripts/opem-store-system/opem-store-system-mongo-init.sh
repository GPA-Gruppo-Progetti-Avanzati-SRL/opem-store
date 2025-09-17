#!/bin/bash

mongosh $1  --file apicore-session.js
mongosh $1  --file apicore-site.js
mongosh $1  --file apicore-domain.js
mongosh $1  --file apicore-user.js
mongosh $1  --file apicore-session.js
mongosh $1  --file apicore-kv.js
mongosh $1  --file apicore-count-documents.js