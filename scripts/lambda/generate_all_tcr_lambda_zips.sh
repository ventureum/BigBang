#!/usr/bin/env bash

echo "generating tcr_table_creator zip"
.//generate_lambda_zip.sh TCR tcr_table_creator $1 $2

echo "generating project zip"
.//generate_lambda_zip.sh TCR project $1 $2

echo "generating get_project zip"
.//generate_lambda_zip.sh TCR get_project $1 $2

echo "generating get_project_list zip"
.//generate_lambda_zip.sh TCR get_project_list $1 $2

echo "generating delete_project zip"
.//generate_lambda_zip.sh TCR delete_project $1 $2

echo "generating objective zip"
.//generate_lambda_zip.sh TCR objective $1 $2

echo "generating get_objective zip"
.//generate_lambda_zip.sh TCR get_objective $1 $2

echo "generating delete_objective zip"
.//generate_lambda_zip.sh TCR delete_objective $1 $2

echo "generating milestone zip"
.//generate_lambda_zip.sh TCR milestone $1 $2

echo "generating get_milestone zip"
.//generate_lambda_zip.sh TCR get_milestone $1 $2

echo "generating delete_milestone zip"
.//generate_lambda_zip.sh TCR delete_milestone $1 $2


echo "generating rating_vote zip"
.//generate_lambda_zip.sh TCR rating_vote $1 $2

echo "generating get_rating_vote_list zip"
.//generate_lambda_zip.sh TCR get_rating_vote_list $1 $2


echo "generating update_actor_rating_votes zip"
.//generate_lambda_zip.sh TCR update_actor_rating_votes $1 $2

#echo "generating add_proxy zip"
#.//generate_lambda_zip.sh TCR add_proxy $1 $2
#
#echo "generating delete_proxy zip"
#.//generate_lambda_zip.sh TCR delete_proxy $1 $2
#
#echo "generating get_proxy_list zip"
#.//generate_lambda_zip.sh TCR get_proxy_list $1 $2
#
#echo "generating rating_vote zip"
#.//generate_lambda_zip.sh TCR rating_vote $1 $2
#

