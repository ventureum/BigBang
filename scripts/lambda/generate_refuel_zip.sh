#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/refuel/
GOOS=linux go build -o main
zip refuel.zip main

#aws lambda create-function \
#    --region ca-central-1 \
#    --function-name refuel_v3_test \
#    --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/refuel/refuel.zip \
#    --role arn:aws:iam::727151012682:role/lambda-vpc-execution-role \
#    --handler main \
#    --runtime go1.x \
#    --environment Variables="{DB_HOST=feedsystest.cmtkgtusnicp.ca-central-1.rds.amazonaws.com,DB_NAME=feedsystest,DB_USER=root,DB_PASSWORD=root1234}"

aws lambda update-function-code \
  --function-name refuel_v3_test \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/refuel/refuel.zip \
  --publish

mv refuel.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
