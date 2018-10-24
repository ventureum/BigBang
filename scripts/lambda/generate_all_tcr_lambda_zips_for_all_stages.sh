#!/usr/bin/env bash

echo "generating test"
.//generate_all_tcr_lambda_zips.sh v3  test

echo "generating staging"
.//generate_all_tcr_lambda_zips.sh v3 staging

echo "generating exp"
.//generate_all_tcr_lambda_zips.sh v3 exp
