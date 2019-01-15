package feed_attributes

import "log"

type VoteType string

const DOWN_VOTE_TYPE VoteType = "DOWN"
const UP_VOTE_TYPE VoteType = "UP"
const LOOKUP_VOTE_TYPE VoteType = "LOOKUP_VOTE_TYPE"

func CreateVoteTypeFromValue(voteValue int64) VoteType {
	var voteType VoteType
	switch voteValue {
	case 1:
		voteType = UP_VOTE_TYPE
	case -1:
		voteType = DOWN_VOTE_TYPE
	case 0:
		voteType = LOOKUP_VOTE_TYPE
	default:
		log.Panicf("Invalid Vote value: %d", voteValue)
	}
	return voteType
}

func (voteType VoteType) Value() int64 {
	var value int64
	switch voteType {
	case UP_VOTE_TYPE:
		value = 1
	case DOWN_VOTE_TYPE:
		value = -1
	default:
		log.Panicf("Invalid Vote Type: %s", voteType)
	}
	return value
}
