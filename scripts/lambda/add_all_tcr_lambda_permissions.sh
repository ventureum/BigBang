#!/usr/bin/env bash

echo "adding lambda permission for new_project"
.//add_tcr_lambda_permission.sh new_project $1 $2 new-project

echo "adding lambda permission for get_project"
.//add_tcr_lambda_permission.sh get_project $1 $2 get-project

echo "adding lambda permission for get_project_list"
.//add_tcr_lambda_permission.sh get_project_list $1 $2 get-project-list

echo "adding lambda permission for add_proxy"
.//add_tcr_lambda_permission.sh add_proxy $1 $2 add-proxy

echo "adding lambda permission for delete_proxy"
.//add_tcr_lambda_permission.sh delete_proxy $1 $2 delete-proxy

echo "adding lambda permission for get_proxy_list"
.//add_tcr_lambda_permission.sh get_proxy_list $1 $2 get-proxy-list

echo "adding lambda permission for rating_vote"
.//add_tcr_lambda_permission.sh rating_vote $1 $2 rating-vote

echo "adding lambda permission for get_rating_vote_list"
.//add_tcr_lambda_permission.sh get_rating_vote_list $1 $2 get-rating-vote-list
