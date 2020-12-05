#!/bin/bash

for i in $(seq 1 $1)
do
echo ""
echo "Request $i"
curl -XPOST -H "Content-type: application/json" -d '{
    "message": "Hello",
    "last_path": "/jump",
    "jump_path": "/jump",
    "jumps": [
        "http://springboot:8443"
    ]
}' 'https://golang-project-a.apps.acidonpe.sandbox1604.opentlc.com/jump' -k 
done
