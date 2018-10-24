#!/usr/bin/env bash

echo "add_all_tcr_lambda_permissions for test"
.//add_all_tcr_lambda_permissions.sh v3 test

echo "add_all_tcr_lambda_permissions for staging"
.//add_all_tcr_lambda_permissions.sh v3 staging

echo "add_all_tcr_lambda_permissions for exp"
.//add_all_tcr_lambda_permissions.sh v3 exp
