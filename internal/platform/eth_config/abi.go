package eth_config

const PostEventABI string = `[{
  "anonymous": false,
  "inputs": [
    {
      "indexed": true,
      "name": "actor",
      "type": "address"
    },
    { 
      "indexed": true,
      "name": "boardId",
      "type": "bytes32"
    },
    {
      "indexed": false,
      "name": "parentHash",
      "type": "bytes32"
    },
    {
      "indexed": true,
      "name": "postHash",
      "type": "bytes32"
    },
    {
      "indexed": false,
      "name": "ipfsPath",
      "type": "bytes32"
    },
    {
      "indexed": false,
      "name": "typeHash",
      "type": "bytes4"
    },
    {
      "indexed": false,
      "name": "timestamp",
      "type": "uint256"
    }
  ],
  "name": "Post",
  "type": "event"
}]`

const UpvoteEventABI = `[{
  "anonymous": false,
  "inputs": [
    {
      "indexed": true,
      "name": "actor",
      "type": "address"
    },
    {
      "indexed": true,
      "name": "boardId",
      "type": "bytes32"
    },
    {
      "indexed": true,
      "name": "postHash",
      "type": "bytes32"
    },
    {
      "indexed": false,
      "name": "value",
      "type": "uint256"
    },
    {
      "indexed": false,
      "name": "timestamp",
      "type": "uint256"
    }
  ],
  "name": "Upvote",
  "type": "event"
}]`

const PurchaseReputationABI = `[{
  "anonymous": false,
  "inputs": [
    {
      "indexed": true,
      "name": "msgSender",
      "type": "address"
    },
    {
      "indexed": true,
      "name": "purchaser",
      "type": "address"
    },
    {
      "indexed": false,
      "name": "numVetX",
      "type": "uint256"
    },
    {
      "indexed": false,
      "name": "numReputation",
      "type": "uint256"
    },
    {
      "indexed": false,
      "name": "timestamp",
      "type": "uint256"
    }
  ],
  "name": "PurchaseReputation",
  "type": "event"
}]`
