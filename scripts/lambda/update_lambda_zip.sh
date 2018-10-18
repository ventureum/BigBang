#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/$1/$2/main/
GOOS=linux go build -o main
zip $2.zip main

DB_NAME=$DB_NAME_PREFIX$4
DB_HOST=$DB_NAME_PREFIX$4$DB_HOST_POSTFIX


if [ "$4" == "staging" ]
then
aws lambda update-function-configuration \
  --function-name $2_$3_$4 \
  --region ca-central-1 \
  --environment Variables="{DEBUG_MODE=0,MAX_FUEL_FOR_FUEL_UPDATE_INTERVAL=$MAX_FUEL_FOR_FUEL_UPDATE_INTERVAL,STREAM_API_KEY=$STREAM_API_KEY_STAGING,STREAM_API_SECRET=$STREAM_API_SECRET_STAGING,FUEL_REPLENISHMENT_HOURLY=$FUEL_REPLENISHMENT_HOURLY,REFUEL_INTERVAL=$REFUEL_INTERVAL,DB_HOST=$DB_HOST,DB_NAME=$DB_NAME,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,MuMaxFuel=$MuMaxFuel,MuMinFuel=$MuMinFuel,PostFuelCost=$PostFuelCost,ReplyFuelCost=$ReplyFuelCost,AuditFuelCost=$AuditFuelCost,BetaMax=$BetaMax}"
else
aws lambda update-function-configuration \
  --function-name $2_$3_$4 \
  --region ca-central-1 \
  --environment Variables="{DEBUG_MODE=0,MAX_FUEL_FOR_FUEL_UPDATE_INTERVAL=$MAX_FUEL_FOR_FUEL_UPDATE_INTERVAL,STREAM_API_KEY=$STREAM_API_KEY,STREAM_API_SECRET=$STREAM_API_SECRET,FUEL_REPLENISHMENT_HOURLY=$FUEL_REPLENISHMENT_HOURLY,REFUEL_INTERVAL=$REFUEL_INTERVAL,DB_HOST=$DB_HOST,DB_NAME=$DB_NAME,DB_USER=$DB_USER,DB_PASSWORD=$DB_PASSWORD,MuMaxFuel=$MuMaxFuel,MuMinFuel=$MuMinFuel,PostFuelCost=$PostFuelCost,ReplyFuelCost=$ReplyFuelCost,AuditFuelCost=$AuditFuelCost,BetaMax=$BetaMax}"
fi

aws lambda update-function-code \
  --function-name $2_$3_$4 \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/$1/$2/main/$2.zip \
  --publish

rm -rf main
rm -rf $2.zip
