#!/bin/bash

usage="$(basename "$0") [-h] [-f] [-p] -- helper script to run simple version of this benchmark on localhost

       where:
           -h  show this help text
           -p  set the server's port (default: 9090)
           -f  set the serialization format to benchmark (acceptable values: json, protobuf, flatbuffers, capnproto)"

PORT=9090
FMT=json

while getopts h:p:f: option
do
    case "${option}"
    in
        h) echo "$usage"
           exit
           ;;
        p) PORT=${OPTARG};;
        f) FMT=${OPTARG};;

    esac
done

cd ../
export ROOTPATH=$(pwd)

if [[ "$OSTYPE" == "linux-gnu" ]]; then
    export FOLDER=linux
elif [[ "$OSTYPE" == "darwin"* ]]; then
    export FOLDER=mac
elif [[ "$OSTYPE" == "cygwin" ]]; then
    export FOLDER=windows
elif [[ "$OSTYPE" == "msys" ]]; then
    export FOLDER=windows
elif [[ "$OSTYPE" == "win32" ]]; then
    export FOLDER=windows
fi


($ROOTPATH/dist/$FOLDER/benchmark_server -p $PORT) &
sleep 5
$ROOTPATH/dist/$FOLDER/benchmark_client -c 10 -d 1000 -details -f $FMT -l 5000 -u "http://localhost:$PORT/data/"
pkill benchmark_server