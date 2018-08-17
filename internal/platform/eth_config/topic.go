package eth_config

import (
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/crypto"
)

var PostEventTopic common.Hash = crypto.Keccak256Hash([]byte("Post(address,bytes32,bytes32,bytes32,bytes32,bytes4,uint256)"))

var UpvoteEventTopic common.Hash = crypto.Keccak256Hash([]byte("Upvote(address,bytes32,bytes32,uint256,uint256)"))

var PurchaseReputationsEventTopic common.Hash = crypto.Keccak256Hash([]byte("PurchaseReputation(address,address,uint256,uint256,uint256)"))
