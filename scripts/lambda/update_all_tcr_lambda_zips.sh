#!/usr/bin/env bash

echo "updating tcr_table_creator zip"
.//update_lambda_zip.sh TCR tcr_table_creator $1 $2

echo "updating new_project zip"
.//update_lambda_zip.sh TCR new_project $1 $2

echo "updating get_project zip"
.//update_lambda_zip.sh TCR get_project $1 $2

echo "updating get_project_list zip"
.//update_lambda_zip.sh TCR get_project_list $1 $2

echo "updating add_proxy zip"
.//update_lambda_zip.sh TCR add_proxy $1 $2

echo "updating delete_proxy zip"
.//update_lambda_zip.sh TCR delete_proxy $1 $2

echo "updating get_proxy_list zip"
.//update_lambda_zip.sh TCR get_proxy_list $1 $2

echo "updating rating_vote zip"
.//update_lambda_zip.sh TCR rating_vote $1 $2

echo "updating get_rating_vote_list zip"
.//update_lambda_zip.sh TCR get_rating_vote_list $1 $2
