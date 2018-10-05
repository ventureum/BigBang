#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/$1/$2/main/
GOOS=linux go build -o main
zip $2.zip main

DB_NAME=$DB_NAME_PREFIX$4
DB_HOST=$DB_NAME_PREFIX$4$DB_HOST_POSTFIX


aws lambda create-function \
    --region ca-central-1 \
    --function-name $2_$3_$4 \
    --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/$1/$2/main/$2.zip \
    --role arn:aws:iam::727151012682:role/lambda-vpc-execution-role \
    --handler main \
    --runtime go1.x \
    --environment Variables="{DEBUG_MODE=0,MAX_FUEL_FOR_FUEL_UPDATE_INTERVAL=$MAX_FUEL_FOR_FUEL_UPDATE_INTERVAL,STREAM_API_KEY=$STREAM_API_KEY,STREAM_API_SECRET=$STREAM_API_SECRET,FUEL_REPLENISHMENT_HOURLY=$FUEL_REPLENISHMENT_HOURLY,REFUEL_INTERVAL=$REFUEL_INTERVAL,DB_HOST=$DB_HOST,DB_NAME=$DB_NAME,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,MuMaxFuel=$MuMaxFuel,MuMinFuel=$MuMinFuel,PostFuelCost=$PostFuelCost,ReplyFuelCost=$ReplyFuelCost,AuditFuelCost=$AuditFuelCost,BetaMax=$BetaMax}"

rm -rf main
rm -rf $2.zip
