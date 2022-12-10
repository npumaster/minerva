#!/bin/sh

Dir=$(cd `dirname $0`; pwd)
cd ${Dir}

echo "build start"

cd ../
go build -o ./build/minerva
cd ${Dir}

cd ../server
go build -o ../build/minerva-server
cd ${Dir}

echo "build end"
