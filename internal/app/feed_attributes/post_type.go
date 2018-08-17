package feed_attributes

import "github.com/ethereum/go-ethereum/crypto"
import "github.com/ethereum/go-ethereum/common"

type PostType string

const (
  PostPostType    PostType = "POST"
  ReplyPostType   PostType = "COMMENT"
  AuditPostType   PostType = "AUDIT"
  AirdropPostType PostType = "AIRDROP"
  SessionPostType    PostType = "SESSION"
)

func (postType PostType) Hash() string {
  bytes4Hash := crypto.Keccak256Hash([]byte(postType.Value())).Bytes()[:4]
  return "0x" + common.Bytes2Hex(bytes4Hash)
}

func (postType PostType) Value() string {
  return string(postType)
}

func CreatePostTypeFromHashStr(typeHashStr string) PostType {
  var postType PostType
  switch typeHashStr {
    case PostPostType.Hash():
      postType = PostPostType
    case ReplyPostType.Hash():
      postType = ReplyPostType
    case AuditPostType.Hash():
      postType = AuditPostType
    case AirdropPostType.Hash():
      postType = AirdropPostType
  }
  return postType
}
