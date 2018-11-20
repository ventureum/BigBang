#!/usr/bin/env bash
bazel test \
    --test_output=errors  \
    --sandbox_debug  \
    --action_env="DEBUG_MODE"\
    --action_env="DEV_ENV" \
    --action_env="DB_USER_TEST" \
    --action_env="DB_PASSWORD_TEST" \
    --action_env="DB_HOST_TEST" \
    --action_env="DB_NAME_TEST" \
    --action_env="STREAM_API_KEY" \
    --action_env="STREAM_API_SECRET" \
    --action_env="MuMaxFuel" \
    --action_env="MuMinFuel" \
    --action_env="PostFuelCost" \
    --action_env="ReplyFuelCost" \
    --action_env="AuditFuelCost" \
    --action_env="BetaMax" \
    --action_env="REFUEL_INTERVAL" \
    --action_env="FUEL_REPLENISHMENT_HOURLY" \
    --action_env="MAX_FUEL_FOR_FUEL_UPDATE_INTERVAL" \
    $1
