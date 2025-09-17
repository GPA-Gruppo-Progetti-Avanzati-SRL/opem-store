#!/bin/bash

mongosh $1  --file apicms-file.js
mongosh $1  --file apicms-count-documents.js