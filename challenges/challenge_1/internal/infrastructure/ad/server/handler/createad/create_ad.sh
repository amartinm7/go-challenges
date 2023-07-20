#bin/bash

curl -X POST -v --location 'http://localhost:8000/v1/ad' \
--header 'Content-Type: application/json' \
--data '{
	"id": "6e96b838-4e11-4409-8035-9aa4ff9e5848",
	"title": "opel astra",
	"description": "como nuevo",
	"price": 15000,
	"timeStamp": "2022-12-01"
}'