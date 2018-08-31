#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/feed_update_post_rewards/
GOOS=linux go build -o main
zip feed_update_post_rewards.zip main

#aws lambda create-function \
#    --region ca-central-1 \
#    --function-name feed_update_post_rewards_v3_test \
#    --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/feed_update_post_rewards/feed_update_post_rewards.zip \
#    --role arn:aws:iam::727151012682:role/lambda-vpc-execution-role \
#    --handler main \
#    --runtime go1.x \
#    --environment Variables="{DB_HOST=feedsystest.cmtkgtusnicp.ca-central-1.rds.amazonaws.com,DB_NAME=feedsystest,DB_USER=root,DB_PASSWORD=root1234,STREAM_API_KEY=6jyjb65k5dxf,STREAM_API_SECRET=csyv2d62k5n6j7femujjb9m8s3md993r8q4tfrjmjvfmt782famuxnehnxuxrrrn}"
#

aws lambda update-function-code \
  --function-name feed_update_post_rewards \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/feed_update_post_rewards/feed_update_post_rewards.zip \
  --publish

mv feed_update_post_rewards.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
