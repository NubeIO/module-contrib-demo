#!/bin/bash

MODULE_NAME="module-contrib-demo"
BIOS_CONTAINER="bios"

docker build -t module-builder -f Dockerfile.module --build-arg="MODULE_NAME=$MODULE_NAME" .
docker run -d --name module-builder module-builder:latest
docker container cp module-builder:/app/$MODULE_NAME .
docker rm -f module-builder
docker exec $BIOS_CONTAINER sh -c 'rm -rf /data/rubix-os/data/modules/'$MODULE_NAME'/; mkdir -p /data/rubix-os/data/modules/'$MODULE_NAME'/v0.0.0/'
docker cp ./$MODULE_NAME $BIOS_CONTAINER:/data/rubix-os/data/modules/$MODULE_NAME/v0.0.0/
docker exec $BIOS_CONTAINER systemctl restart nubeio-rubix-os.service
