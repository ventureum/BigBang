#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/get_session/
GOOS=linux go build -o main
zip get_session.zip main

#aws lambda create-function \
#    --region ca-central-1 \
#    --function-name get_session_v3_test \
#    --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/get_session/get_session.zip \
#    --role arn:aws:iam::727151012682:role/lambda-vpc-execution-role \
#    --handler main \
#    --runtime go1.x \
#    --environment Variables="{DB_HOST=feedsystest.cmtkgtusnicp.ca-central-1.rds.amazonaws.com,DB_NAME=feedsystest,DB_USER=root,DB_PASSWORD=root1234}"

aws lambda update-function-code \
  --function-name get_session_v3_test \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/get_session/get_session.zip \
  --publish

mv get_session.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
