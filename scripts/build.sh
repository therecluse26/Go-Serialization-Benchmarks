#!/bin/bash

if [[ "$OSTYPE" == "linux-gnu" ]]; then
    export GOOS=linux
    export FOLDER=linux
elif [[ "$OSTYPE" == "darwin"* ]]; then
    export GOOS=darwin
    export FOLDER=mac
elif [[ "$OSTYPE" == "cygwin" || "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    export GOOS=windows
    export FOLDER=windows
fi

if [[ $(uname -m) == "x86_64" ]]; then
    export GOARCH=amd64
elif [[ $(uname -m) == "i386" || $(uname -m) == "i686" ]]; then
    export GOARCH=386
elif [[ $(uname -m) == "aarch64" ]]; then
    export GOARCH=arm64
fi

cd ../
export ROOTPATH=$(pwd)

cd $ROOTPATH/cmd/client
go build -o $ROOTPATH/dist/$FOLDER/benchmark_client
cd $ROOTPATH/cmd/server
go build -o $ROOTPATH/dist/$FOLDER/benchmark_server