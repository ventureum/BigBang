#!/usr/bin/env bash

aws apigateway update-method \
    --rest-api-id $1 \
    --resource-id $2 \
    --http-method POST \
    --patch-operations "[{\"op\": \"replace\", \"path\": \"/authorizationType\", \"value\": \"CUSTOM\"}, {\"op\": \"replace\", \"path\": \"/authorizerId\", \"value\": \"$4\"}]"

aws apigateway put-integration \
    --rest-api-id $1 \
    --resource-id $2 \
    --http-method POST \
    --integration-http-method POST \
    --type AWS \
    --uri "arn:aws:apigateway:ca-central-1:lambda:path/2015-03-31/functions/arn:aws:lambda:ca-central-1:727151012682:function:$3\${stageVariables.postfix} /invocations" \
    --passthrough-behavior WHEN_NO_TEMPLATES

aws apigateway update-integration \
    --rest-api-id $1 \
    --resource-id $2 \
    --http-method POST \
    --patch-operations "op='add',path='/requestTemplates/application~1json',value='{\"principalId\": \"\$context.authorizer.principalId\",\"body\": \$input.json(\'\$\')}'"
