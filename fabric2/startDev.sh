#!/usr/bin/env bash


./network.sh down
rm -rf organizations/peerOrganizations
rm -rf organizations/ordererOrganizations
rm -rf organizations/rest-certs
docker network create cc-mensalist-net
./network.sh up createChannel
./network.sh deployCC -ccn cc-mensalist -ccp ../chaincode -ccl go