#!/usr/bin/env bash

echo "updating test"
.//update_all_tcr_lambda_zips.sh v3  test

echo "updating staging"
.//update_all_tcr_lambda_zips.sh v3 staging

echo "updating exp"
.//update_all_tcr_lambda_zips.sh v3 exp
