#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/profile/
GOOS=linux go build -o main
zip profile.zip main


#aws lambda create-function \
#    --region ca-central-1 \
#    --function-name profile_v3_test \
#    --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/profile/profile.zip \
#    --role arn:aws:iam::727151012682:role/lambda-vpc-execution-role \
#    --handler main \
#    --runtime go1.x \
#    --environment Variables="{DB_HOST=feedsystest.cmtkgtusnicp.ca-central-1.rds.amazonaws.com,DB_NAME=feedsystest,DB_USER=root,DB_PASSWORD=root1234}"
#

#aws lambda update-function-configuration \
#  --function-name profile_v3_test \
#  --region ca-central-1 \
#  --environment Variables="{DB_HOST=$DB_HOST,DB_NAME=$DB_NAME,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,MuMaxFuel=$MuMaxFuel,MuMinFuel=$MuMinFuel,PostFuelCost=$PostFuelCost,ReplyFuelCost=$ReplyFuelCost,AuditFuelCost=$AuditFuelCost,BetaMax=$BetaMax}"

aws lambda update-function-code \
  --function-name profile_v3_test \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/profile/profile.zip \
  --publish

mv profile.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
