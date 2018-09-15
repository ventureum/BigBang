#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/$1/
GOOS=linux go build -o main
zip $1.zip main

DB_NAME=$DB_NAME_PREFIX$3
DB_HOST=$DB_NAME_PREFIX$2$DB_HOST_POSTFIX


aws lambda create-function \
    --region ca-central-1 \
    --function-name $1_$2_$3 \
    --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/$1/$1.zip \
    --role arn:aws:iam::727151012682:role/lambda-vpc-execution-role \
    --handler main \
    --runtime go1.x \
    --environment Variables="{FUEL_REPLENISHMENT_HOURLY=$FUEL_REPLENISHMENT_HOURLY,REFUEL_INTERVAL=$REFUEL_INTERVAL,DB_HOST=$DB_HOST,DB_NAME=$DB_NAME,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,MuMaxFuel=$MuMaxFuel,MuMinFuel=$MuMinFuel,PostFuelCost=$PostFuelCost,ReplyFuelCost=$ReplyFuelCost,AuditFuelCost=$AuditFuelCost,BetaMax=$BetaMax}"

rm -rf main
rm -rf $1.zip
