package eth_config

import (
  "os"
  "log"
  "context"

  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/accounts/abi"
  "strings"
  "reflect"
  "time"
  "errors"
  "BigBang/internal/app/feed_attributes"
  "BigBang/internal/platform/postgres_config/client_config"
  "BigBang/internal/platform/postgres_config/post_config"
  "BigBang/internal/platform/postgres_config/post_votes_record_config"
  "BigBang/internal/platform/postgres_config/purchase_reputations_record_config"
  "BigBang/internal/platform/postgres_config/actor_reputations_record_config"
  "BigBang/internal/platform/postgres_config/post_replies_record_config"
  "BigBang/internal/platform/postgres_config/post_reputations_record_config"
  "BigBang/internal/platform/getstream_config"
)


type EthClient struct {
  c *ethclient.Client
}


const LOCAL_SOCKET_URL string = "ws://127.0.0.1:8546"


func ConnectEthClient() (*EthClient) {
  rawURL := os.Getenv("SOCKET_URL")

  if rawURL == "" {
    rawURL = LOCAL_SOCKET_URL
  }

  client, err := ethclient.Dial(rawURL)
  if err != nil {
    log.Panicf("Failed to dial eth client with url %s with error: %+v\n", rawURL, err)
  }
  log.Println("Connected to Ethereum EthClient")
  return &EthClient{c:  client}
}

func (client *EthClient) Close() {
  client.c.Close()
  log.Println("Disconnected to Ethereum EthClient")
}

func createFilterQuery(forumAddressHex string) ethereum.FilterQuery {
  forumAddress := common.HexToAddress(forumAddressHex)
  query := ethereum.FilterQuery{
    Addresses: []common.Address{forumAddress},
    Topics: [][]common.Hash{{
      PostEventTopic,
      UpvoteEventTopic,
      PurchaseReputationsEventTopic,
    }},
  }
  return query
}

func (client *EthClient) SubscribeFilterLogs(
    forumAddressHex string, getStreamClient *getstream_config.GetStreamClient, postgresFeedClient *client_config.PostgresFeedClient) {
  logs := make(chan types.Log)
  filterQuery := createFilterQuery(forumAddressHex)
  sub, err := client.c.SubscribeFilterLogs(context.Background(), filterQuery, logs)
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Subscribed to FilterLogs")
  for {
    select {
      case err := <-sub.Err():
        log.Printf("SubscribeFilterLogs Error: %+v", err)
      case vLog := <-logs:
         err := ProcessRequest(vLog, getStreamClient, postgresFeedClient)
         if err != nil {
           log.Printf("Failed to process Log %+v with error: %+v\n", vLog, err)
         }
    }
  }
}

func ProcessRequest(
    vLog types.Log,
    getStreamClient *getstream_config.GetStreamClient,
    postgresFeedClient *client_config.PostgresFeedClient) (err error) {
  defer func() {
    if errStr := recover(); errStr != nil { //catch
      err = errors.New(errStr.(string))
    }
  }()
  event, err := matchEvent(vLog.Topics, vLog.Data)
  if err != nil {
    log.Panicf("Error to match event: %+v", err)
  }
  log.Printf("Processing Event: %+v\n", *event)
  processEvent(event, getStreamClient, postgresFeedClient)
  return err
}

func matchEvent(topics []common.Hash, data []byte) (*Event, error) {
  if len(topics) == 0 {
    return nil, nil
  }
  var event Event
  switch topics[0] {
    case PostEventTopic:
      var postEventResult PostEventResult
      postEventAbi, _ := abi.JSON(strings.NewReader(PostEventABI))
      err := postEventAbi.Unpack(&postEventResult, "Post", data)
      if err != nil {
        return nil, err
      }
      postEventResult.Poster = common.BytesToAddress(topics[1].Bytes())
      postEventResult.BoardId = topics[2]
      postEventResult.PostHash = topics[3]
      event = *postEventResult.ToPostRecord()
      return &event, nil

    case UpvoteEventTopic:
      var upvoteEventResult UpvoteEventResult
      upvoteEventAbi, _ := abi.JSON(strings.NewReader(UpvoteEventABI))
      err := upvoteEventAbi.Unpack(&upvoteEventResult, "Upvote", data)
      if err != nil {
        return nil, err
      }
      upvoteEventResult.Actor = common.BytesToAddress(topics[1].Bytes())
      upvoteEventResult.BoardId = topics[2]
      upvoteEventResult.PostHash = topics[3]
      event = *upvoteEventResult.ToPostVotesRecord()
      return &event, nil

    case PurchaseReputationsEventTopic:
      var purchaseReputationsEventResult PurchaseReputationsEventResult
      purchaseReputationsEventAbi, _ := abi.JSON(strings.NewReader(PurchaseReputationABI))
      err := purchaseReputationsEventAbi.Unpack(&purchaseReputationsEventResult, "PurchaseReputation", data)
      if err != nil {
        return nil, err
      }
      purchaseReputationsEventResult.MsgSender = common.BytesToAddress(topics[1].Bytes())
      purchaseReputationsEventResult.Purchaser = common.BytesToAddress(topics[2].Bytes())
      event = *purchaseReputationsEventResult.ToPurchaseReputationsRecord()
      return &event, nil
  }

  return nil, nil
}

func processEvent(
    event *Event,
    getStreamClient *getstream_config.GetStreamClient,
    postgresFeedClient *client_config.PostgresFeedClient) {
  switch reflect.TypeOf(*event) {
    case reflect.TypeOf(post_config.PostRecord{}):
      postRecord := (*event).(post_config.PostRecord)
      ProcessPostRecord(&postRecord, getStreamClient, postgresFeedClient, feed_attributes.ON_CHAIN)
    case reflect.TypeOf(post_votes_record_config.PostVotesRecord{}):
      postVotesRecord := (*event).(post_votes_record_config.PostVotesRecord)
      ProcessPostVotesRecord(&postVotesRecord, postgresFeedClient)
    case reflect.TypeOf(purchase_reputations_record_config.PurchaseReputationsRecord{}):
      purchaseReputationsRecord := (*event).(purchase_reputations_record_config.PurchaseReputationsRecord)
      ProcessPurchaseReputationsRecord(&purchaseReputationsRecord, postgresFeedClient)
  }
}

func ProcessPostRecord(
    postRecord *post_config.PostRecord,
    getStreamClient *getstream_config.GetStreamClient,
    postgresFeedClient *client_config.PostgresFeedClient,
    source feed_attributes.Source) {
  postgresFeedClient.Begin()
  postExecutor := post_config.PostExecutor{*postgresFeedClient}
  actorReputationsRecordExecutor := actor_reputations_record_config.ActorReputationsRecordExecutor{
    *postgresFeedClient}
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresFeedClient}

  updateCount :=  postExecutor.GetPostUpdateCountTx(postRecord.PostHash)
  reputationsPenalty := feed_attributes.PenaltyForPostType(
    feed_attributes.PostType(postRecord.PostType), updateCount)

  // Update Actor Reputation
  actorReputationsRecordExecutor.SubActorReputationsTx(postRecord.Actor, reputationsPenalty)

  // Insert Post Record
  updatedTimestamp := postExecutor.UpsertPostRecordTx(postRecord)

  // Insert Activity to GetStream
  activity := ConvertPostRecordToActivity(postRecord, source, feed_attributes.BlockTimestamp(updatedTimestamp.Unix()))
  getStreamClient.AddFeedActivityToGetStream(activity)

  // Update Post Replies Record
  if activity.Verb == feed_attributes.ReplyVerb {
    postRepliesRecord := post_replies_record_config.PostRepliesRecord {
      PostHash: postRecord.ParentHash,
      ReplyHash: postRecord.PostHash,
    }
    postRepliesRecordExecutor.UpsertPostRepliesRecordTx(&postRepliesRecord)
  }

  postgresFeedClient.Commit()
}

func ProcessPostVotesRecord(
    postVotesRecord *post_votes_record_config.PostVotesRecord,
    postgresFeedClient *client_config.PostgresFeedClient) {
  postgresFeedClient.Begin()

  actorReputationsRecordExecutor := actor_reputations_record_config.ActorReputationsRecordExecutor{
    *postgresFeedClient}
  postReputationsRecordExecutor := post_reputations_record_config.PostReputationsRecordExecutor{*postgresFeedClient}
  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*postgresFeedClient}

  // CutOff Time
  cutOffTimeStamp := time.Now()

  // Actor List for PostHash and VoteType
  actorList := *postReputationsRecordExecutor.GetActorListByPostHashAndVoteTypeTx(
    postVotesRecord.PostHash, postVotesRecord.VoteType)

  log.Printf("Actor List for PostHash and VoteType: %+v\n", actorList)

  // Current Actor Reputation
  actorReputation := actorReputationsRecordExecutor.GetActorReputationsTx(postVotesRecord.Actor)

  log.Printf("Current Actor Reputation: %+v\n", actorReputation)

  // Total Actor Reputations
  totalActorReputations := actorReputationsRecordExecutor.GetTotalActorReputationsTx()

  log.Printf("Total Actor Reputations: %+v\n", totalActorReputations)

  // Total Actor Reputations for PostHash
  totalReputationsForPostHash := postReputationsRecordExecutor.GetTotalReputationsByPostHashTx(postVotesRecord.PostHash)


  log.Printf("Total Actor Reputations for PostHash: %+v\n", totalReputationsForPostHash)

  // Total Actor Reputations for PostHash with the same voteType as actor
  totalReputationsForPostHashWithSameVoteType := postReputationsRecordExecutor.GetReputationsByPostHashAndVoteTypeTx(
    postVotesRecord.PostHash, postVotesRecord.VoteType)

  log.Printf("Total Actor Reputations for PostHash with the same voteType as actor: %+v\n",
    totalReputationsForPostHashWithSameVoteType)

  // Last Actor Reputation when doing vote
  lastActorReputation := postReputationsRecordExecutor.GetReputationsByPostHashAndActorTx(
    postVotesRecord.PostHash, postVotesRecord.Actor)

  log.Printf("Last Actor Reputation when doing vote: %+v\n", lastActorReputation)

  totalReputationsForPostHash  =  totalReputationsForPostHash - lastActorReputation + actorReputation
  totalReputationsForPostHashWithSameVoteType = totalReputationsForPostHashWithSameVoteType -
    lastActorReputation + actorReputation

  log.Printf("Updated  totalReputationsForPostHash: %+v\n",  totalReputationsForPostHash)
  log.Printf("Updated  totalReputationsForPostHashWithSameVoteType : %+v\n",
    totalReputationsForPostHashWithSameVoteType)
  log.Printf("totalActorReputations : %+v\n",  totalActorReputations )

  // Calculate Vote Cost
  voteCost := post_votes_record_config.STACK_FRACTION * float64(actorReputation)*
      (1.00 - float64(totalReputationsForPostHash) / float64(totalActorReputations))

  log.Printf("voteCost: %+v\n", voteCost)

  voteCount :=  postReputationsRecordExecutor.GetTotalVotesCountByPostHashAndActorType(
    postVotesRecord.PostHash, postVotesRecord.Actor)

  log.Printf("Vote Count: %+v\n", voteCount)

  votePenalty := feed_attributes.PenaltyForVote(feed_attributes.Reputation(voteCost), voteCount)

  log.Printf("vote Penalty : %+v\n", votePenalty)

  // Deduct  votePenalty
  actorReputationsRecordExecutor.SubActorReputationsTx(postVotesRecord.Actor, votePenalty)

  // Record current vote
  postVotesRecord.SignedReputations = actorReputation.Value() * postVotesRecord.VoteType.Value()
  postVotesRecordExecutor.UpsertPostVotesRecordTx(postVotesRecord)

  // Update Actor Reputation For the postHash
  postReputationsRecord := post_reputations_record_config.PostReputationsRecord{
    Actor: postVotesRecord.Actor,
    PostHash: postVotesRecord.PostHash,
    Reputations: actorReputation.SubReputations(votePenalty),
    LatestVoteType: postVotesRecord.VoteType,
  }
  postReputationsRecordExecutor.UpsertPostReputationsRecordTx(&postReputationsRecord)


  if totalReputationsForPostHashWithSameVoteType > 0 {
    // Distribute Rewards
    for _, actorAddress := range actorList {
      awardedActorReputation := postReputationsRecordExecutor.GetReputationsByPostHashAndActorWithLatestVoteTypeAndTimeCutOffTx(
          postVotesRecord.PostHash,
          actorAddress,
          postVotesRecord.VoteType,
          cutOffTimeStamp)
      rewards :=  int64(float64(votePenalty) * float64(awardedActorReputation) / float64(totalReputationsForPostHashWithSameVoteType))

      log.Printf(" rewards  %+v for actorAddress %s\n",  rewards , actorAddress)
      actorReputationsRecordExecutor.AddActorReputationsTx(actorAddress, feed_attributes.Reputation(rewards))
    }
  }

  postgresFeedClient.Commit()
}

func ProcessPurchaseReputationsRecord(
    purchaseReputationsRecord *purchase_reputations_record_config.PurchaseReputationsRecord,
    postgresFeedClient *client_config.PostgresFeedClient) {
  postgresFeedClient.Begin()

  actorReputationsRecordExecutor := actor_reputations_record_config.ActorReputationsRecordExecutor{
    *postgresFeedClient}
  purchaseReputationsRecordExecutor := purchase_reputations_record_config.PurchaseReputationsRecordExecutor{
    *postgresFeedClient}
  purchaseReputationsRecordExecutor.UpsertPurchaseReputationsRecordTx(purchaseReputationsRecord)
  actorReputationsRecordExecutor.AddActorReputationsTx(
    purchaseReputationsRecord.Purchaser, feed_attributes.Reputation(purchaseReputationsRecord.Reputations))

  postgresFeedClient.Commit()
}

func ConvertPostRecordToActivity(
    postRecord *post_config.PostRecord,
    source feed_attributes.Source,
    timestamp feed_attributes.BlockTimestamp) *feed_attributes.Activity {
  var verb feed_attributes.Verb
  var to []feed_attributes.FeedId
  var obj feed_attributes.Object
  extraParam := map[string]interface{}{
    "source": source,
  }

  if postRecord.ParentHash == NullHashString {
    obj = feed_attributes.Object{
      ObjType:feed_attributes.PostObjectType,
      ObjId: postRecord.PostHash,
    }
    verb = feed_attributes.SubmitVerb
    to = []feed_attributes.FeedId {
      {
        FeedSlug: feed_attributes.BoardFeedSlug,
        UserId: feed_attributes.AllBoardId,
      },
      {
        FeedSlug: feed_attributes.BoardFeedSlug,
        UserId: feed_attributes.UserId(postRecord.BoardId),
      },
    }
  } else {
    obj = feed_attributes.Object{
      ObjType:feed_attributes.ReplyObjectType,
      ObjId: postRecord.PostHash,
    }
    verb = feed_attributes.ReplyVerb
    extraParam["post"] = feed_attributes.Object{
        ObjType: feed_attributes.PostObjectType,
        ObjId: postRecord.ParentHash,
    }
    to = []feed_attributes.FeedId {
      {
        FeedSlug: feed_attributes.CommentFeedSlug,
        UserId: feed_attributes.UserId(postRecord.ParentHash),
      },
    }
  }

  actor := feed_attributes.Actor(postRecord.Actor)
  postType := feed_attributes.PostType(postRecord.PostType)
  return feed_attributes.CreateNewActivity(actor, verb, obj, timestamp, postType, to, extraParam)
}
