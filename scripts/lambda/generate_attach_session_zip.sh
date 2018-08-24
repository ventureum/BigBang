#!/usr/bin/env bash

cd $GOPATH/src/BigBang/cmd/lambda/attach_session/
GOOS=linux go build -o main
zip attach_session.zip main

aws lambda update-function-code \
  --function-name attach_session \
  --zip-file fileb://$GOPATH/src/BigBang/cmd/lambda/attach_session/attach_session.zip \
  --publish

mv attach_session.zip $GOPATH/src/BigBang/assets/lambda_zips
rm -rf main
