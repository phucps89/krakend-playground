#!/bin/bash
rootDir="${0%/*}"
$rootDir/krakend run -c "$rootDir/krakend.json"