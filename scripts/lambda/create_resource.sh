#!/usr/bin/env bash

aws apigateway create-resource \
  --rest-api-id $1 \
  --parent-id $2 \
  --path-part $3

