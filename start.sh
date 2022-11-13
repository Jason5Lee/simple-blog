#!/bin/bash

pushd "$(dirname "$0")"
docker compose up
popd
