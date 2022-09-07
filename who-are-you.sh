#! /bin/bash
curl -S https://ytrack.learn.ynov.com/assets/superhero/all.json | jq '.[] | select(.id == 70) .name'