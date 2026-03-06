#!/bin/bash

#curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq -r --arg HERO_ID $HERO_ID ' .[] | select(.id==($HERO_ID|tonumber)) | .connections.relatives' | sed ':a;N;$!ba;s/\n/\\n/g';
#!/bin/bash

allTheData=$(curl -s https://acad.learn2earn.ng/assets/superhero/all.json)


id=$(jq -r --arg id "$HERO_ID" ' .[] | select(.id==($id|tonumber)) | .connections.relatives' <<< "$allTheData")

printf "$id" | sed ':a;N;$!ba;s/\n/\\n/g';