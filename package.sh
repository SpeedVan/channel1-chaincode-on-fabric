#!/bin/sh

export FABRIC_CFG_PATH=/home/szy/projects/github.com/SpeedVan/supply-chain-finance-on-fabric/base/fabric/config

go mod vendor
peer lifecycle chaincode package channel1.tar.gz --path . --lang golang --label c_2.0

cp channel1.tar.gz /home/szy/projects/github.com/SpeedVan/supply-chain-finance-on-fabric/gyl_core_org1/channel1-chaincode