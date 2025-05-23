#!/bin/bash
if [ -z ${DOCKER_HUB_USER+x} ]
then
    echo 'Skipping login - credentials not set'
else
    docker login -u $DOCKER_HUB_USER -p $DOCKER_HUB_PASSWORD
fi

docker-compose -f server/docker-compose.yaml push