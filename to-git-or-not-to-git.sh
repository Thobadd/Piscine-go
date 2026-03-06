#!/bin/bash


allTheData=$(curl -s https://acad.learn2earn.ng/assets/superhero/all.json)

name=$(jq ' .[] | select(.id==170) | .name' <<< "$allTheData")

printf "$name"

power=$(jq ' .[] | select(.id==170) | .powerstats.power' <<< "$allTheData")

printf "$power"

gender=$(jq ' .[] | select(.id==170) | .appearance.gender' <<<"$allTheData")

printf "$gender"
