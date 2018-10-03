#!/usr/bin/env bash

echo "updating testing"
.//update_lambda_zip.sh $1 $2 test

echo "updating staging"
.//update_lambda_zip.sh $1 $2 staging

echo "updating production"
.//update_lambda_zip.sh $1 $2 producction
