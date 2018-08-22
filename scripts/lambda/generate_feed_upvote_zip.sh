#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/feed_upvote/
GOOS=linux go build -o main
zip feed_upvote.zip main

aws lambda update-function-code \
  --function-name feed_upvote_v2 \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/feed_upvote/feed_upvote.zip \
  --publish

mv feed_upvote.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
