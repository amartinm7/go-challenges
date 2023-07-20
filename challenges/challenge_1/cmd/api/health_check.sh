#bin/bash

curl -X GET -v --location 'http://localhost:8000/health' --header 'Content-Type: application/json'
