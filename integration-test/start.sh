#!/bin/bash

pushd "$(dirname "$0")"
docker compose -p simple-blog-integration-test up
popd
