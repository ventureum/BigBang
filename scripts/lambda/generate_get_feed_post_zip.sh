#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/get_feed_post/
GOOS=linux go build -o main
zip get_feed_post.zip main


#aws lambda create-function \
#    --region ca-central-1 \
#    --function-name get_feed_post_v3_test \
#    --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/get_feed_post/get_feed_post.zip \
#    --role arn:aws:iam::727151012682:role/lambda-vpc-execution-role \
#    --handler main \
#    --runtime go1.x \
#    --environment Variables="{DB_HOST=feedsystest.cmtkgtusnicp.ca-central-1.rds.amazonaws.com,DB_NAME=feedsystest,DB_USER=root,DB_PASSWORD=root1234,STREAM_API_KEY=6jyjb65k5dxf,STREAM_API_SECRET=csyv2d62k5n6j7femujjb9m8s3md993r8q4tfrjmjvfmt782famuxnehnxuxrrrn}"

#aws lambda update-function-configuration \
#  --function-name get_feed_post_v3_test \
#  --region ca-central-1 \
#  --environment Variables="{STREAM_API_KEY=$STREAM_API_KEY,STREAM_API_SECRET=$STREAM_API_SECRET,DB_HOST=$DB_HOST,DB_NAME=$DB_NAME,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,MuMaxFuel=$MuMaxFuel,MuMinFuel=$MuMinFuel,PostFuelCost=$PostFuelCost,ReplyFuelCost=$ReplyFuelCost,AuditFuelCost=$AuditFuelCost,BetaMax=$BetaMax}"
#

aws lambda update-function-code \
  --function-name get_feed_post_v3_test \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/get_feed_post/get_feed_post.zip \
  --publish

mv get_feed_post.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
