package eth_config

import (
  "github.com/ethereum/go-ethereum/common"
  "BigBang/internal/pkg/utils"
)

var PostEventTopic common.Hash = utils.Keccak256Hash([]byte("Post(address,bytes32,bytes32,bytes32,bytes32,bytes4,uint256)"))

var UpvoteEventTopic common.Hash = utils.Keccak256Hash([]byte("Upvote(address,bytes32,bytes32,uint256,uint256)"))

var PurchaseReputationsEventTopic common.Hash = utils.Keccak256Hash([]byte("PurchaseReputation(address,address,uint256,uint256,uint256)"))
