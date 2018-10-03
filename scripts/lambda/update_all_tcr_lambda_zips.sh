#!/usr/bin/env bash

echo "updating tcr_table_creator zip"
.//update_lambda_zip.sh TCR tcr_table_creator $1 $2

echo "updating new_project zip"
.//update_lambda_zip.sh TCR new_project $1 $2

echo "updating get_project zip"
.//update_lambda_zip.sh TCR get_project $1 $2

echo "updating get_project_list zip"
.//update_lambda_zip.sh TCR get_project_list $1 $2
