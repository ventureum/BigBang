#!/usr/bin/env bash

rm -rf $GOPATH/src/BIGBANG/third_party/libsecp256k1
mkdir -p $GOPATH/src/BIGBANG/third_party/libsecp256k1
mkdir -p $GOPATH/src/BIGBANG/third_party/libsecp256k1/include
mkdir -p $GOPATH/src/BIGBANG/third_party/libsecp256k1/lib

cd $GOPATH/src/BIGBANG/vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1
./autogen.sh
./configure
make
./tests
cp include/*.h  $GOPATH/src/BIGBANG/third_party/libsecp256k1/include
cp src/*.h $GOPATH/src/BIGBANG/third_party/libsecp256k1/include
cp src/modules/recovery/main_impl.h $GOPATH/src/BIGBANG/third_party/libsecp256k1/include
cp .libs/libsecp256k1.a $GOPATH/src/BIGBANG/third_party/libsecp256k1/lib
cp .libs/libsecp256k1.so $GOPATH/src/BIGBANG/third_party/libsecp256k1/lib