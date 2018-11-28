#!/usr/bin/env bash

TCR_REST_API_ID=mfmybdhoea
FEED_REST_API_ID=7g1vjuevub
FEED_AUTHORIZER=3my87y
TCR_AUTHORIZER=w3apiu
TEST_API_ID=8hpbrf
Lambda_FUNCTION_NAME=add_addresses

echo "set up auth for test api"
.//set_up_auth.sh $FEED_REST_API_ID $TEST_API_ID Lambda_FUNCTION_NAME $FEED_AUTHORIZER

.//set_up_auth.sh $TCR_REST_API_ID  ynwk6p delete_batch_objectives $TCR_AUTHORIZER
.//set_up_auth.sh $TCR_REST_API_ID  yfpezm delete_batch_objectives $TCR_AUTHORIZER