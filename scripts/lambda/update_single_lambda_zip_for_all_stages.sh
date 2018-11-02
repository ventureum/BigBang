#!/usr/bin/env bash

.//update_lambda_zip.sh $1 $2 v3 test

.//update_lambda_zip.sh $1 $2 v3 staging

.//update_lambda_zip.sh $1 $2 v3 exp