#bin/bash

curl -v --location 'http://localhost:8000/health' \
--header 'Content-Type: application/json'
