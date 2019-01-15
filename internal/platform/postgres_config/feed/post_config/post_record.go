package post_config

import (
	"BigBang/internal/app/feed_attributes"
	"BigBang/internal/pkg/utils"
	"github.com/jmoiron/sqlx/types"
	"time"
)

type PostRecord struct {
	Actor       string
	BoardId     string         `db:"board_id"`
	ParentHash  string         `db:"parent_hash"`
	PostHash    string         `db:"post_hash"`
	PostType    string         `db:"post_type"`
	Content     types.JSONText `db:"content"`
	UpdateCount int64          `db:"update_count"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}

type PostRecordResult struct {
	Actor       string
	BoardId     string
	ParentHash  string
	PostHash    string
	PostType    string
	Content     *feed_attributes.Content
	UpdateCount int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (postRecord *PostRecord) ToPostRecordResult() *PostRecordResult {
	return &PostRecordResult{
		Actor:       postRecord.Actor,
		BoardId:     postRecord.BoardId,
		ParentHash:  postRecord.ParentHash,
		PostHash:    postRecord.PostHash,
		PostType:    postRecord.PostType,
		Content:     feed_attributes.CreatedContentFromJsonText(postRecord.Content),
		UpdateCount: postRecord.UpdateCount,
		CreatedAt:   postRecord.CreatedAt,
		UpdatedAt:   postRecord.UpdatedAt,
	}
}

func (postRecord *PostRecord) ToActivity(
	source feed_attributes.Source,
	timestamp feed_attributes.BlockTimestamp) *feed_attributes.Activity {
	var verb feed_attributes.Verb
	var to []feed_attributes.FeedId
	var obj feed_attributes.Object
	extraParam := map[string]interface{}{
		"source": source,
	}

	if postRecord.ParentHash == utils.NullHashString {
		obj = feed_attributes.Object{
			ObjType: feed_attributes.PostObjectType,
			ObjId:   postRecord.PostHash,
		}
		verb = feed_attributes.SubmitVerb
		to = []feed_attributes.FeedId{
			{
				FeedSlug: feed_attributes.BoardFeedSlug,
				UserId:   feed_attributes.AllBoardId,
			},
			{
				FeedSlug: feed_attributes.BoardFeedSlug,
				UserId:   feed_attributes.UserId(postRecord.BoardId),
			},
		}
	} else {
		obj = feed_attributes.Object{
			ObjType: feed_attributes.ReplyObjectType,
			ObjId:   postRecord.PostHash,
		}
		verb = feed_attributes.ReplyVerb
		extraParam["post"] = feed_attributes.Object{
			ObjType: feed_attributes.PostObjectType,
			ObjId:   postRecord.ParentHash,
		}
		to = []feed_attributes.FeedId{
			{
				FeedSlug: feed_attributes.CommentFeedSlug,
				UserId:   feed_attributes.UserId(postRecord.ParentHash),
			},
		}
	}

	actor := feed_attributes.Actor(postRecord.Actor)
	postType := feed_attributes.PostType(postRecord.PostType)
	return feed_attributes.CreateNewActivity(actor, verb, obj, timestamp, postType, to, extraParam)
}
