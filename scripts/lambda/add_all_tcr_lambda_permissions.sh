#!/usr/bin/env bash

echo "adding lambda permission for project"
.//add_tcr_lambda_permission.sh project $1 $2 project

echo "adding lambda permission for get_project"
.//add_tcr_lambda_permission.sh get_project $1 $2 get-project

echo "adding lambda permission for get_project_list"
.//add_tcr_lambda_permission.sh get_project_list $1 $2 get-project-list

echo "adding lambda permission for delete_project"
.//add_tcr_lambda_permission.sh delete_project $1 $2 delete-project


echo "adding lambda permission for milestone"
.//add_tcr_lambda_permission.sh milestone $1 $2 milestone

echo "adding lambda permission for get_milestone"
.//add_tcr_lambda_permission.sh get_milestone $1 $2 get-milestone

echo "adding lambda permission for delete_milestone"
.//add_tcr_lambda_permission.sh delete_milestone $1 $2 delete-milestone

echo "adding lambda permission for objective"
.//add_tcr_lambda_permission.sh objective $1 $2 objective

echo "adding lambda permission for get_objective"
.//add_tcr_lambda_permission.sh get_objective $1 $2 get-objective

echo "adding lambda permission for delete_objective"
.//add_tcr_lambda_permission.sh delete_objective $1 $2 delete-objective

echo "adding lambda permission for rating_vote"
.//add_tcr_lambda_permission.sh rating_vote $1 $2 rating-vote

echo "adding lambda permission for get_rating_vote_list"
.//add_tcr_lambda_permission.sh get_rating_vote_list $1 $2 get-rating-vote-list

echo "adding lambda permission for update_actor_rating_votes"
.//add_tcr_lambda_permission.sh update_actor_rating_votes $1 $2 update-actor-rating-votes

echo "adding lambda permission for adjust_proxy_votes"
.//add_tcr_lambda_permission.sh adjust_proxy_votes $1 $2 adjust-proxy-votes

#echo "adding lambda permission for add_proxy"
#.//add_tcr_lambda_permission.sh add_proxy $1 $2 add-proxy
#
#echo "adding lambda permission for delete_proxy"
#.//add_tcr_lambda_permission.sh delete_proxy $1 $2 delete-proxy
#
#echo "adding lambda permission for get_proxy_list"
#.//add_tcr_lambda_permission.sh get_proxy_list $1 $2 get-proxy-list
#
#echo "adding lambda permission for rating_vote"
#.//add_tcr_lambda_permission.sh rating_vote $1 $2 rating-vote
#
#echo "adding lambda permission for get_rating_vote_list"
#.//add_tcr_lambda_permission.sh get_rating_vote_list $1 $2 get-rating-vote-list
