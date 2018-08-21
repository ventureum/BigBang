#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/feed_post/
GOOS=linux go build -o main
zip feed_post.zip main

aws lambda update-function-code \
  --function-name feed_post_v2 \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/feed_post/feed_post.zip \
  --publish

mv feed_post.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
