#!/bin/bash

code_path=$1

if [ "$code_path" = "" ];
then
    code_path="code/go"
    echo NO PATH PROVIDED WILL USE: $code_path
else
    echo PATH PROVIDED WILL USE: $code_path
fi

path=$HOME/$code_path/rubix-os

echo $path
echo "****EXISTING MODULES****"
ls  $path/data/modules
echo "****EXISTING MODULES****"

go build -o module-contrib-demo && mkdir -p $path/data/modules && mv -f module-contrib-demo $path/data/modules && cd $path && bash build.bash
