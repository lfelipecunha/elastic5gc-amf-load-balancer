#!/bin/bash
scp -P 8282 -r "$1" soft5g@ppgca.unisinos.br:"~/load-balancer/$(dirname $1)"
