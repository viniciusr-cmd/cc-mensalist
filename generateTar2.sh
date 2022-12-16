#!/usr/bin/env bash

rm -f cc-mensalist.tar.gz

# Make sure go mod is up to date
cd chaincode && go mod vendor && cd ..

export FABRIC_CFG_PATH=fabric2/config
peer lifecycle chaincode package chaincode.tar.gz --path chaincode --lang golang --label cc-mensalist_1.0

# Compress file without rest-server (GoFabric will use the standard CC API)
tar -czf cc-mensalist.tar.gz chaincode.tar.gz

# Compress file with rest-server (GoFabric will use the one provided)
# tar -c --exclude=node_modules -zf cc-mensalist.tar.gz chaincode.tar.gz rest-server

rm -f chaincode.tar.gz