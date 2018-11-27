#!/bin/bash

rm app.zip
zip -r app.zip ./
aws lambda update-function-code \
    --region "ca-central-1" \
    --function-name "subscribe_boards_$1_$2" \
    --zip-file "fileb://./app.zip"
