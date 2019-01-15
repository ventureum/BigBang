package feed_attributes

import (
	"BigBang/internal/pkg/utils"
	"log"
	"strings"
)

type FeedSlug string
type UserId string
type FeedId struct {
	FeedSlug FeedSlug `json:"feedSlug"`
	UserId   UserId   `json:"userId"`
}

const (
	UserFeedSlug        FeedSlug = "user"
	UserPostFeedSlug    FeedSlug = "user_post"
	BoardFeedSlug       FeedSlug = "board"
	UserCommentFeedSlug FeedSlug = "user_comment"
	CommentFeedSlug     FeedSlug = "comment"
)

var AllBoardId = UserId(utils.Keccak256Hash([]byte("%__AllBoardId__%")).String())

func CreateFeedId(feedSlug string, userid string) FeedId {
	return FeedId{
		FeedSlug: FeedSlug(feedSlug),
		UserId:   UserId(userid),
	}
}

func CreateFeedIdFromValue(value string) FeedId {
	s := strings.Split(value, ":")
	if len(s) != 2 {
		log.Fatal("value is not valid when creating Feed Id: ", value)
	}
	return CreateFeedId(s[0], s[1])
}

func (feedId FeedId) Value() string {
	return string(feedId.FeedSlug) + ":" + string(feedId.UserId)
}

func (feedId FeedId) FeedToken(secret string) string {
	id := strings.Replace(feedId.Value(), ":", "", -1)
	return utils.CryptoToken(id, secret)
}

func ConvertToStringArray(feedIds []FeedId) []string {
	arr := make([]string, len(feedIds))
	for i, v := range feedIds {
		arr[i] = v.Value()
	}
	return arr
}

func ConvertFromStringArrayToFeedIds(feedIds []string) []FeedId {
	arr := make([]FeedId, len(feedIds))
	for i, v := range feedIds {
		arr[i] = CreateFeedIdFromValue(v)
	}
	return arr
}
