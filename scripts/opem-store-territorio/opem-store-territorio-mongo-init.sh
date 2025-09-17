#!/bin/bash

mongosh $1  --file opem-store-territorio-nazione.js
mongosh $1  --file opem-store-territorio-provincia.js
mongosh $1  --file opem-store-territorio-comune.js
mongosh $1  --file opem-store-territorio-count-documents.js