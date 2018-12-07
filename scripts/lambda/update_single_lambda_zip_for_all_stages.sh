#!/usr/bin/env bash

.//update_lambda_zip.sh $1 $2 v3 test $4

.//update_lambda_zip.sh $1 $2 v3 staging $4

.//update_lambda_zip.sh $1 $2 v3 exp $4