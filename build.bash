#!/bin/bash

path=$1

echo "***Path***"
echo $path
echo "****List module****"
ls  $path/data/rubix-os/modules/module-contrib-demo/v0.0.0/

go build -o module-contrib-demo && mkdir -p $path/data/rubix-os/modules/module-contrib-demo/v0.0.0/ && mv -f module-contrib-demo $path/data/rubix-os/modules/module-contrib-demo/v0.0.0/ && cd $path && ./app-amd64
