#!/usr/bin/env bash

#echo "generating tcr_table_creator zip"
#.//generate_lambda_zip.sh TCR tcr_table_creator $1 $2
#
#echo "generating new_project zip"
#.//generate_lambda_zip.sh TCR new_project $1 $2
#
#echo "generating get_project zip"
#.//generate_lambda_zip.sh TCR get_project $1 $2
#
#echo "generating get_project_list zip"
#.//generate_lambda_zip.sh TCR get_project_list $1 $2
#
#echo "generating get_project_list zip"
#.//generate_lambda_zip.sh TCR get_project_list $1 $2

echo "generating add_proxy zip"
.//generate_lambda_zip.sh TCR add_proxy $1 $2

echo "generating delete_proxy zip"
.//generate_lambda_zip.sh TCR delete_proxy $1 $2

echo "generating get_proxy_list zip"
.//generate_lambda_zip.sh TCR get_proxy_list $1 $2
