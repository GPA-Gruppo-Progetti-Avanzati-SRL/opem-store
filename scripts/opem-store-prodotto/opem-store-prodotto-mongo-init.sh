#!/bin/bash

mongosh $1  --file opem-store-prodotto-prodotto.js
mongosh $1  --file opem-store-prodotto-count-documents.js