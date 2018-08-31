#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/feed_events_table_creator/
GOOS=linux go build -o main
zip feed_events_table_creator.zip main

#aws lambda create-function \
#    --region ca-central-1 \
#    --function-name feed_events_table_creator_v3_test \
#    --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/feed_events_table_creator/feed_events_table_creator.zip \
#    --role arn:aws:iam::727151012682:role/lambda-vpc-execution-role \
#    --handler main \
#    --runtime go1.x \
#    --environment Variables="{STREAM_API_KEY=$STREAM_API_KEY,STREAM_API_SECRET=$STREAM_API_SECRET,DB_HOST=$DB_HOST,DB_NAME=$DB_NAME,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,MuMaxFuel=$MuMaxFuel,MuMinFuel=$MuMinFuel,PostFuelCost=$PostFuelCost,ReplyFuelCost=$ReplyFuelCost,AuditFuelCost=$AuditFuelCost,BetaMax=$BetaMax}"

aws lambda update-function-code \
  --function-name feed_events_table_creator_v3_test \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/feed_events_table_creator/feed_events_table_creator.zip \
  --publish

mv feed_events_table_creator.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
