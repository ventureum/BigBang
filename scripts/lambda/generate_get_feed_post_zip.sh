#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/get_feed_post/
GOOS=linux go build -o main
zip get_feed_post.zip main

aws lambda update-function-code \
  --function-name get_feed_post_v2 \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/get_feed_post/get_feed_post.zip \
  --publish

mv get_feed_post.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
