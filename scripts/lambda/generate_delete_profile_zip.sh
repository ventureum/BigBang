#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/delete_profile/
GOOS=linux go build -o main
zip delete_profile.zip main

aws lambda update-function-code \
  --function-name delete_profile \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/delete_profile/delete_profile.zip \
  --publish

mv delete_profile.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
