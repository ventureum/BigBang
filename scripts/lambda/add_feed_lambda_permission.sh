#!/usr/bin/env bash

FUNCTION_NAME=$1
FUNCTION=$1_$2_$3

ID=0

aws lambda remove-permission \
   --function-name "arn:aws:lambda:ca-central-1:727151012682:function:$FUNCTION"   \
   --statement-id $(($ID+1))

aws lambda add-permission   \
   --function-name "arn:aws:lambda:ca-central-1:727151012682:function:$FUNCTION"   \
   --source-arn "arn:aws:execute-api:ca-central-1:727151012682:7g1vjuevub/*/POST/$4"   \
   --principal apigateway.amazonaws.com \
   --statement-id $(($ID+1)) \
   --action lambda:InvokeFunction

aws lambda remove-permission \
   --function-name "arn:aws:lambda:ca-central-1:727151012682:function:$FUNCTION"   \
   --statement-id $(($ID+2))

aws lambda add-permission   \
   --function-name "arn:aws:lambda:ca-central-1:727151012682:function:$FUNCTION"   \
   --source-arn "arn:aws:execute-api:ca-central-1:727151012682:7g1vjuevub/*/OPTIONS/$4"   \
   --principal apigateway.amazonaws.com \
   --statement-id $(($ID+2)) \
   --action lambda:InvokeFunction
