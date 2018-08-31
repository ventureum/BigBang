#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/attach_session/
GOOS=linux go build -o main
zip attach_session.zip main

#aws lambda create-function \
#    --region ca-central-1 \
#    --function-name attach_session_v3_test \
#    --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/attach_session/attach_session.zip \
#    --role arn:aws:iam::727151012682:role/lambda-vpc-execution-role \
#    --handler main \
#    --runtime go1.x \
#    --environment Variables="{STREAM_API_KEY=$STREAM_API_KEY,STREAM_API_SECRET=$STREAM_API_SECRET,DB_HOST=$DB_HOST,DB_NAME=$DB_NAME,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,MuMaxFuel=$MuMaxFuel,MuMinFuel=$MuMinFuel,PostFuelCost=$PostFuelCost,ReplyFuelCost=$ReplyFuelCost,AuditFuelCost=$AuditFuelCost,BetaMax=$BetaMax}"
#

aws lambda update-function-code \
  --function-name attach_session_v3_test \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/attach_session/attach_session.zip \
  --publish

mv attach_session.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
