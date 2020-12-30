#!/bin/sh

export FABRIC_CFG_PATH=/Users/admin/projects/go/src/github.com/hyperledger/fabric-samples/config

go mod vendor
peer lifecycle chaincode package basic.tar.gz --path . --lang golang --label basic_1.0