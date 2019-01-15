package feed_attributes

import (
	"BigBang/internal/pkg/error_config"
	"BigBang/internal/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

type PostType string

const (
	PostPostType          PostType = "POST"
	ReplyPostType         PostType = "COMMENT"
	RatingCommentPostType PostType = "RATING_COMMENT"
	AuditPostType         PostType = "AUDIT"
	AirdropPostType       PostType = "AIRDROP"
)

func (postType PostType) Hash() string {
	bytes4Hash := utils.Keccak256Hash([]byte(postType.Value())).Bytes()[:4]
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
	case RatingCommentPostType.Hash():
		postType = AirdropPostType
	default:
		errorInfo := error_config.ErrorInfo{
			ErrorCode: error_config.InvalidPostType,
			ErrorData: map[string]interface{}{
				"typeHash": typeHashStr,
			},
			ErrorLocation: error_config.PostTypeLocation,
		}
		log.Printf("Invalid typeHash: %s", typeHashStr)
		log.Panicln(errorInfo.Marshal())
	}
	return postType
}
