package post_votes_record_config

import (
  "log"
  "BigBang/internal/platform/postgres_config/post_votes_record_config"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/client_config"
)


func main() {
  db := client_config.ConnectPostgresClient()
  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*db}
  postVotesRecordExecutor.LoadVoteTypeEnum()
  postVotesRecordExecutor.DeletePostVotesRecordTable()
  postVotesRecordExecutor.CreatePostVotesRecordTable()

  actorOne := "0x01"
  actorTwo := "0x02"
  postHashOne := "0xpostHash001"
  postHashTwo := "0xpostHash002"
  postHashThree := "0xpostHash003"

  postVotesRecord1 := &post_votes_record_config.PostVotesRecord{
    Actor: actorOne,
    PostHash: postHashOne,
    VoteType: feed_attributes.UP_VOTE_TYPE,
  }

  postVotesRecord2 := &post_votes_record_config.PostVotesRecord{
    Actor: actorOne,
    PostHash: postHashOne,
    VoteType: feed_attributes.DOWN_VOTE_TYPE,
  }

  postVotesRecord3 := &post_votes_record_config.PostVotesRecord{
    Actor: actorOne,
    PostHash: postHashOne,
    VoteType: feed_attributes.DOWN_VOTE_TYPE,
  }

  postVotesRecord4 := &post_votes_record_config.PostVotesRecord{
    Actor: actorTwo,
    PostHash: postHashOne,
    VoteType: feed_attributes.UP_VOTE_TYPE,
  }

  postVotesRecord5 := &post_votes_record_config.PostVotesRecord{
    Actor: actorTwo,
    PostHash: postHashTwo,
    VoteType: feed_attributes.DOWN_VOTE_TYPE,
  }

  postVotesRecord6 := &post_votes_record_config.PostVotesRecord{
    Actor: actorTwo,
    PostHash: postHashTwo,
    VoteType: feed_attributes.DOWN_VOTE_TYPE,
  }

  postVotesRecordExecutor.UpsertPostVotesRecord(postVotesRecord1)
  postVotesRecordExecutor.UpsertPostVotesRecord(postVotesRecord2)
  postVotesRecordExecutor.UpsertPostVotesRecord(postVotesRecord3)
  postVotesRecordExecutor.UpsertPostVotesRecord(postVotesRecord4)
  postVotesRecordExecutor.UpsertPostVotesRecord(postVotesRecord5)
  postVotesRecordExecutor.UpsertPostVotesRecord(postVotesRecord6)

  postVotesForActorOneAndPostHashOne := postVotesRecordExecutor.GetTotalPostVotesCount(
    actorOne, postHashOne)
  log.Printf("total votes for actor %s and postHash %s: %+v\n",
    actorOne, postHashOne, postVotesForActorOneAndPostHashOne)

  upPostVotesForActorOneAndPostHashOne := postVotesRecordExecutor.GetVotesCountByVoteType(
    actorOne, postHashOne, feed_attributes.UP_VOTE_TYPE)
  log.Printf("up votes for actor %s and postHash %s: %+v\n",
    actorOne, postHashOne, upPostVotesForActorOneAndPostHashOne)

  downPostVotesForActorOneAndPostHashOne := postVotesRecordExecutor.GetVotesCountByVoteType(
    actorOne, postHashOne,  feed_attributes.DOWN_VOTE_TYPE)
  log.Printf("down votes for actor %s and postHash %s: %+v\n",
    actorOne, postHashOne, downPostVotesForActorOneAndPostHashOne)

  postVotesForActorTwoAndPostHashOne := postVotesRecordExecutor.GetTotalPostVotesCount(
    actorTwo, postHashOne)
  log.Printf("total votes for actor %s and postHash %s: %+v\n",
    actorTwo, postHashOne, postVotesForActorTwoAndPostHashOne)

  upPostVotesForActorTwoAndPostHashOne := postVotesRecordExecutor.GetVotesCountByVoteType(
    actorTwo, postHashOne, feed_attributes.UP_VOTE_TYPE)
  log.Printf("up votes for actor %s and postHash %s: %+v\n",
    actorTwo, postHashOne, upPostVotesForActorTwoAndPostHashOne)

  downPostVotesForActorTwoAndPostHashOne := postVotesRecordExecutor.GetVotesCountByVoteType(
    actorTwo, postHashOne, feed_attributes.DOWN_VOTE_TYPE)
  log.Printf("down votes for actor %s and postHash %s: %+v\n",
    actorTwo, postHashOne, downPostVotesForActorTwoAndPostHashOne)

  postVotesForActorTwoAndPostHashTwo := postVotesRecordExecutor.GetTotalPostVotesCount(
    actorTwo, postHashTwo)
  log.Printf("total votes for actor %s and postHash %s: %+v\n",
    actorTwo, postHashTwo, postVotesForActorTwoAndPostHashTwo)

  upPostVotesForActorTwoAndPostHashTwo := postVotesRecordExecutor.GetVotesCountByVoteType(
    actorTwo, postHashTwo, feed_attributes.UP_VOTE_TYPE)
  log.Printf("up votes for actor %s and postHash %s: %+v\n",
    actorTwo, postHashTwo, upPostVotesForActorTwoAndPostHashTwo)

  downPostVotesForActorTwoAndPostHashTwo := postVotesRecordExecutor.GetVotesCountByVoteType(
    actorTwo, postHashTwo, feed_attributes.DOWN_VOTE_TYPE)
  log.Printf("down votes for actor %s and postHash %s: %+v\n",
    actorTwo, postHashTwo, downPostVotesForActorTwoAndPostHashTwo)

  downPostVotesForActorThreeAndPostHashTwo := postVotesRecordExecutor.GetVotesCountByVoteType(
    actorTwo, postHashThree, feed_attributes.DOWN_VOTE_TYPE)
  log.Printf("down votes for actor %s and postHash %s: %+v\n",
    actorTwo, postHashThree, downPostVotesForActorThreeAndPostHashTwo)
}
