#!/usr/bin/env bash

cd ~/challenge/test-network

./network.sh deployCC -c mychannel -ccn assetcc -ccp ../chaincode -ccl go

source ./scripts/setOrgPeerContext.sh 1
export FABRIC_CFG_PATH=${PWD}/configtx

peer chaincode invoke -o localhost:7050 \
--ordererTLSHostnameOverride orderer.example.com \
--tls --cafile $ORDERER_CA -C mychannel -n assetcc \
--peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
--peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
-c '{"Args":["DeleteAsset","asset1"]}'