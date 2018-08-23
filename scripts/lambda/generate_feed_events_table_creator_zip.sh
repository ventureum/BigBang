#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/feed_events_table_creator/
GOOS=linux go build -o main
zip feed_events_table_creator.zip main

aws lambda update-function-code \
  --function-name create_feed_tables \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/feed_events_table_creator/feed_events_table_creator.zip \
  --publish

mv feed_events_table_creator.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
