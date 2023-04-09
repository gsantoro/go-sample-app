#!/usr/bin/env bash

env_var=$1
default_value=$2

if [[ -n "$env_var" ]]; then
    echo "$env_var"
else
    echo "$default_value"
fi