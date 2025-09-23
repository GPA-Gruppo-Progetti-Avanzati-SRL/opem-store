#!/bin/bash

mongosh $1  --file opem-store-magazzino-magazzini.js
mongosh $1  --file opem-store-magazzino-boxes.js
mongosh $1  --file opem-store-magazzino-count-documents.js