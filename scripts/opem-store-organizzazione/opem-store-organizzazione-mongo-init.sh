#!/bin/bash

mongosh $1  --file opem-store-organizzazione-focalpoint.js
mongosh $1  --file opem-store-organizzazione-count-documents.js