#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/$1/
GOOS=linux go build -o main
zip $1.zip main

DB_NAME=$DB_NAME_PREFIX$2
DB_HOST=$DB_NAME_PREFIX$2$DB_HOST_POSTFIX

aws lambda update-function-configuration \
  --function-name $1_$3_$2 \
  --region ca-central-1 \
  --environment Variables="{DB_HOST=$DB_HOST,DB_NAME=$DB_NAME,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,MuMaxFuel=$MuMaxFuel,MuMinFuel=$MuMinFuel,PostFuelCost=$PostFuelCost,ReplyFuelCost=$ReplyFuelCost,AuditFuelCost=$AuditFuelCost,BetaMax=$BetaMax}"

aws lambda update-function-code \
  --function-name $1_$3_$2 \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/$1/$1.zip \
  --publish

rm -rf main
rm -rf $1.zip
