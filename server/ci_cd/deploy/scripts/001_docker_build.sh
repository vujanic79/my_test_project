#!/bin/bash
export BUILD_NUMBER=$1
docker-compose -f server/docker-compose.yaml build