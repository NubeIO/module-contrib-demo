#!/bin/bash

code_path=$1

if [ "$code_path" = "" ];
then
    code_path="code/go"
    echo NO PATH PROVIDED WILL USE: $code_path
else
    echo PATH PROVIDED WILL USE: $code_path
fi

path=$HOME/$code_path/rubix-os-build

echo "****START-BIOS****"

cd $path/rubix-bios && ./rubix-bios server

