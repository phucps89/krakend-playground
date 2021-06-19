#!/bin/bash
rootDir="${0%/*}"
cd $rootDir/krakend-module && \
make build && \
mv krakend ../dist/