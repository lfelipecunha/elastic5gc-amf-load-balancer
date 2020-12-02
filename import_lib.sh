#!/bin/bash

basepath=$(realpath $(dirname $0))

liburl="https://github.com/free5gc/$1.git"

cd "$basepath/lib" && git clone $liburl

rm -rf $basepath/lib/$1/.git
cd "$basepath/lib/$1" && find -type f -exec sed -i 's/free5gc/amfLoadBalancer/g' {} +
