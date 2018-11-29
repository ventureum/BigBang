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
  "BigBang/internal/platform/postgres_config/feed/post_config"
  "BigBang/internal/platform/postgres_config/feed/post_votes_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_replies_record_config"
  "BigBang/internal/platform/getstream_config"
  "BigBang/internal/platform/postgres_config/feed/post_votes_counters_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_profile_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_rewards_info_record_config"
  "BigBang/internal/platform/postgres_config/feed/purchase_mps_record_config"
  "BigBang/internal/platform/postgres_config/feed/actor_votes_counters_record_config"
  "BigBang/internal/platform/postgres_config/feed/post_rewards_record_config"
  "math"
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
    forumAddressHex string, getStreamClient *getstream_config.GetStreamClient, postgresBigBangClient *client_config.PostgresBigBangClient) {
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
         err := ProcessRequest(vLog, getStreamClient, postgresBigBangClient)
         if err != nil {
           log.Printf("Failed to process Log %+v with error: %+v\n", vLog, err)
         }
    }
  }
}

func ProcessRequest(
    vLog types.Log,
    getStreamClient *getstream_config.GetStreamClient,
    postgresBigBangClient *client_config.PostgresBigBangClient) (err error) {
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
  processEvent(event, getStreamClient, postgresBigBangClient)
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
    postgresBigBangClient *client_config.PostgresBigBangClient) {
  switch reflect.TypeOf(*event) {
    case reflect.TypeOf(post_config.PostRecord{}):
      postRecord := (*event).(post_config.PostRecord)
      postgresBigBangClient.Begin()
      ProcessPostRecord(&postRecord, getStreamClient, postgresBigBangClient, feed_attributes.ON_CHAIN)
      postgresBigBangClient.Commit()
    case reflect.TypeOf(post_votes_record_config.PostVotesRecord{}):
      postVotesRecord := (*event).(post_votes_record_config.PostVotesRecord)
      postgresBigBangClient.Begin()
      ProcessPostVotesRecord(&postVotesRecord, postgresBigBangClient)
      postgresBigBangClient.Commit()
    case reflect.TypeOf(purchase_mps_record_config.PurchaseMPsRecord{}):
      purchaseReputationsRecord := (*event).(purchase_mps_record_config.PurchaseMPsRecord)
      ProcessPurchaseReputationsRecord(&purchaseReputationsRecord, postgresBigBangClient)
  }
}

func ProcessPostRecord(
    postRecord *post_config.PostRecord,
    getStreamClient *getstream_config.GetStreamClient,
    postgresBigBangClient *client_config.PostgresBigBangClient,
    source feed_attributes.Source) {
  postExecutor := post_config.PostExecutor{*postgresBigBangClient}
  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
    *postgresBigBangClient}
  postRepliesRecordExecutor := post_replies_record_config.PostRepliesRecordExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  postRewardsRecordExecutor := post_rewards_record_config.PostRewardsRecordExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor.VerifyActorExistingTx(postRecord.Actor)
  actorRewardsInfoRecordExecutor.VerifyActorExistingTx(postRecord.Actor)

  updateCount :=  postExecutor.GetPostUpdateCountTx(postRecord.PostHash)
  fuelsPenalty := feed_attributes.FuelsPenaltyForPostType(
    feed_attributes.PostType(postRecord.PostType), updateCount)

  log.Printf("UpdateCount for PostHash %s: %d", postRecord.PostHash, updateCount)
  log.Printf("Fuel Penalty for PostHash %s: %d", postRecord.PostHash, fuelsPenalty)

  // Update Actor Fuel
  actorRewardsInfoRecordExecutor.SubActorFuelTx(postRecord.Actor, fuelsPenalty)

  // Insert Post Record
  createdTimestamp := postExecutor.UpsertPostRecordTx(postRecord)
  activity := ConvertPostRecordToActivity(postRecord, source, feed_attributes.BlockTimestamp(createdTimestamp.Unix()))

  postRewardsRecordExecutor.UpsertPostRewardsRecordTx(&post_rewards_record_config.PostRewardsRecord{
    PostHash: postRecord.PostHash,
    Actor: postRecord.Actor,
    PostType: postRecord.PostType,
    Object: activity.Object.Value(),
    PostTime: createdTimestamp,
    DeltaFuel: int64(fuelsPenalty.Neg()),
  })

  // Insert Activity to GetStream
  if updateCount == 0 {
    getStreamClient.AddFeedActivityToGetStream(activity)
  } else {
    getStreamClient.UpdateFeedActivityToGetStream(activity)
  }


  // Update Post Replies Record
  if activity.Verb == feed_attributes.ReplyVerb {
    postRepliesRecord := post_replies_record_config.PostRepliesRecord {
      PostHash: postRecord.ParentHash,
      ReplyHash: postRecord.PostHash,
    }
    postRepliesRecordExecutor.UpsertPostRepliesRecordTx(&postRepliesRecord)
  }
}

func ProcessPostVotesRecord(
    postVotesRecord *post_votes_record_config.PostVotesRecord,
    postgresBigBangClient *client_config.PostgresBigBangClient) *feed_attributes.VoteInfo {

  actor := postVotesRecord.Actor
  postHash := postVotesRecord.PostHash
  voteType := postVotesRecord.VoteType

  var voteInfo feed_attributes.VoteInfo
  voteInfo.Actor = postVotesRecord.Actor
  voteInfo.PostHash = postVotesRecord.PostHash

  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
    *postgresBigBangClient}
  actorVotesCountersRecordExecutor := actor_votes_counters_record_config.ActorVotesCountersRecordExecutor{*postgresBigBangClient}
  postVotesRecordExecutor := post_votes_record_config.PostVotesRecordExecutor{*postgresBigBangClient}
  postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  postExecutor := post_config.PostExecutor{*postgresBigBangClient}

  actorProfileRecordExecutor.VerifyActorExistingTx(postVotesRecord.Actor)
  actorRewardsInfoRecordExecutor.VerifyActorExistingTx(postVotesRecord.Actor)
  postExecutor.VerifyPostRecordExistingTx(postVotesRecord.PostHash)

  // CutOff Time
  cutOffTimeStamp := time.Now()

  // Actor List for PostHash and VoteType

  var actorList []string

  actorList = *postVotesRecordExecutor.GetActorListByPostHashAndVoteTypeTx(
    postVotesRecord.PostHash, postVotesRecord.VoteType)
  log.Printf("Actor List for PostHash and VoteType: %+v\n", actorList)


  // Current Actor ActualMilestonePoints
  actorRewardsInfo := actorRewardsInfoRecordExecutor.GetActorRewardsInfoTx(actor)
  log.Printf("Current Actor RewardsInfo: %+v\n", actorRewardsInfo)

  // Total Actor Reputation
  totalActorReputations := actorRewardsInfoRecordExecutor.GetTotalActorReputationTx()
  log.Printf("Total Actor Reputation: %+v\n", totalActorReputations)

  postVotesCountersRecord := postVotesCountersRecordExecutor.GetPostVotesCountersRecordByPostHashTx(postHash)

  // Total Actor Reputation for PostHash
  totalReputationsForPostHash := actorVotesCountersRecordExecutor.GetTotalActorReputationByPostHashTx(postHash)
  log.Printf("Total Actor Reputation for PostHash %s: %+v\n", postHash, totalReputationsForPostHash)


  // Total Reputation for PostHash with the same voteType as actor
  var totalReputationsForPostHashWithSameVoteType feed_attributes.Reputation
  if voteType == feed_attributes.UP_VOTE_TYPE {
    totalReputationsForPostHashWithSameVoteType = postVotesCountersRecord.TotalReputationForUpvote
  } else {
    totalReputationsForPostHashWithSameVoteType = postVotesCountersRecord.TotalReputationForDownvote
  }
  log.Printf("Total Actor Reputaions for PostHash with the same voteType as actor: %+v\n",
    totalReputationsForPostHashWithSameVoteType)

  // Calculate FuelCost
  var fuelCost feed_attributes.Fuel
  if totalActorReputations > 0 {
    fuelCost = feed_attributes.Fuel(math.Round(float64(feed_attributes.BetaMax) * (1.00 - float64(totalReputationsForPostHash)/(float64(totalActorReputations)))))
  } else {
    fuelCost = feed_attributes.BetaMax
  }
  voteInfo.FuelCost = feed_attributes.Fuel(fuelCost)
  log.Printf("FuelCost for PostHash %s: %+v\n", postHash, voteInfo.FuelCost)


  // Update Fuel
  actorRewardsInfoRecordExecutor.SubActorFuelTx(actor, feed_attributes.Fuel(fuelCost))


  // Update Actor Reputation For the postHash
  actorVotesCountersRecord := actor_votes_counters_record_config.ActorVotesCountersRecord{
    Actor:          actor,
    PostHash:       postHash,
    LatestReputation:  actorRewardsInfo.Reputation,
    LatestVoteType: voteType,
  }
  upsertedPostReputationsRecord := actorVotesCountersRecordExecutor.UpsertActorVotesCountersRecordTx(&actorVotesCountersRecord)

  // Record current vote
  postVotesRecord.DeltaFuel = int64(fuelCost.Neg())
  postVotesRecord.DeltaReputation = 0
  postVotesRecord.DeltaMilestonePoints = 0
  postVotesRecord.SignedReputation = actorRewardsInfo.Reputation.Value() * postVotesRecord.VoteType.Value()
  postVotesRecord.PostType = string(postExecutor.GetPostTypeTx(postHash))
  postVotesRecordExecutor.UpsertPostVotesRecordTx(postVotesRecord)

  newPostVotesCountersRecord := post_votes_counters_record_config.PostVotesCountersRecord{
    PostHash: postHash,
    LatestVoteType: voteType,
    LatestActorReputation: actorRewardsInfo.Reputation,
  }
  upsertPostVotesCountersRecord := postVotesCountersRecordExecutor.UpsertPostVotesCountersRecordTx(
    &newPostVotesCountersRecord)

  voteInfo.RewardsInfo = actorRewardsInfoRecordExecutor.GetActorRewardsInfoTx(actor)

  voteInfo.PostVoteCountInfo = &feed_attributes.VoteCountInfo{
    UpVoteCount: upsertPostVotesCountersRecord.UpVoteCount,
    DownVoteCount: upsertPostVotesCountersRecord.DownVoteCount,
    TotalVoteCount: upsertPostVotesCountersRecord.TotalVoteCount,
  }

  voteInfo.RequestorVoteCountInfo = &feed_attributes.VoteCountInfo{
    UpVoteCount: upsertedPostReputationsRecord.UpVoteCount,
    DownVoteCount: upsertedPostReputationsRecord.DownVoteCount,
    TotalVoteCount: upsertedPostReputationsRecord.TotalVoteCount,
  }

  if totalReputationsForPostHashWithSameVoteType > 0 {
    // Distribute Rewards
    for _, actorAddress := range actorList {
      awardedActorReputation := actorVotesCountersRecordExecutor.GetReputationByPostHashAndActorWithLatestVoteTypeAndTimeCutOffTx(
        postVotesRecord.PostHash,
        actorAddress,
        voteType,
        cutOffTimeStamp)
      rewards := int64(float64(fuelCost) * float64(awardedActorReputation) / float64(totalReputationsForPostHashWithSameVoteType))

      log.Printf("rewards %+v for actorAddress %s\n", rewards, actorAddress)
      actorRewardsInfoRecordExecutor.AddActorReputationTx(actorAddress, feed_attributes.Reputation(rewards))
      actorRewardsInfoRecordExecutor.AddActorMilestonePointsFromVotesTx(actorAddress, feed_attributes.MilestonePoint(rewards))
      postVotesRecordExecutor.AddPostVoteDeltaRewardsInfoTx(actorAddress, postHash, voteType, 0, int64(rewards), int64(rewards))
    }
  }

  return &voteInfo
}

func QueryPostVotesInfo(
    postVotesRecord *post_votes_record_config.PostVotesRecord,
    postgresBigBangClient *client_config.PostgresBigBangClient) *feed_attributes.VoteInfo {
  var voteInfo feed_attributes.VoteInfo

  actor := postVotesRecord.Actor
  postHash := postVotesRecord.PostHash

  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
    *postgresBigBangClient}
  actorVotesCountersRecordExecutor := actor_votes_counters_record_config.ActorVotesCountersRecordExecutor{*postgresBigBangClient}
  actorProfileRecordExecutor := actor_profile_record_config.ActorProfileRecordExecutor{*postgresBigBangClient}
  postExecutor := post_config.PostExecutor{*postgresBigBangClient}
  postVotesCountersRecordExecutor := post_votes_counters_record_config.PostVotesCountersRecordExecutor{*postgresBigBangClient}

  actorProfileRecordExecutor.VerifyActorExisting(postVotesRecord.Actor)
  actorRewardsInfoRecordExecutor.VerifyActorExisting(postVotesRecord.Actor)
  postExecutor.VerifyPostRecordExisting(postVotesRecord.PostHash)


  // Current Actor ActualMilestonePoints
  actorRewardsInfo := actorRewardsInfoRecordExecutor.GetActorRewardsInfo(actor)
  log.Printf("Current Actor RewardsInfo: %+v\n", actorRewardsInfo)
  voteInfo.RewardsInfo = actorRewardsInfo

  // Total Actor Reputation
  totalActorReputations := actorRewardsInfoRecordExecutor.GetTotalActorReputation()

  log.Printf("Total Actor Reputation: %+v\n", totalActorReputations)

  postVotesCountersRecord := postVotesCountersRecordExecutor.GetPostVotesCountersRecordByPostHash(postHash)

  // Total Actor Reputation for PostHash
  totalReputationsForPostHash := actorVotesCountersRecordExecutor.GetTotalActorReputationByPostHash(postHash)

  log.Printf("Total Actor Reputation for PostHash %s: %+v\n", postHash, totalReputationsForPostHash)

  // Calculate FuelCost
  var fuelCost feed_attributes.Fuel
  if totalActorReputations > 0 {
    fuelCost = feed_attributes.Fuel(math.Round(float64(feed_attributes.BetaMax) * (1.00 - float64(totalReputationsForPostHash)/(float64(totalActorReputations)))))
  } else {
    fuelCost = feed_attributes.BetaMax
  }
  voteInfo.FuelCost = feed_attributes.Fuel(fuelCost)
  log.Printf("FuelCost for PostHash %s: %+v\n", postHash, voteInfo.FuelCost)

  actorVotesCountersRecord := actorVotesCountersRecordExecutor.GetActorVotesCountersRecordByPostHashAndActor(
    postHash, actor)

  voteInfo.PostHash = postVotesRecord.PostHash
  voteInfo.Actor = postVotesRecord.Actor
  voteInfo.PostVoteCountInfo = &feed_attributes.VoteCountInfo{
    UpVoteCount: postVotesCountersRecord.UpVoteCount,
    DownVoteCount: postVotesCountersRecord.DownVoteCount,
    TotalVoteCount: postVotesCountersRecord.TotalVoteCount,
  }
  voteInfo.RequestorVoteCountInfo = &feed_attributes.VoteCountInfo{
    UpVoteCount:  actorVotesCountersRecord.UpVoteCount,
    DownVoteCount:  actorVotesCountersRecord.DownVoteCount,
    TotalVoteCount:  actorVotesCountersRecord.TotalVoteCount,
  }

  return &voteInfo
}

func ProcessPurchaseReputationsRecord(
    purchaseMPsRecord *purchase_mps_record_config.PurchaseMPsRecord,
    postgresBigBangClient *client_config.PostgresBigBangClient) {
  postgresBigBangClient.Begin()

  actorRewardsInfoRecordExecutor := actor_rewards_info_record_config.ActorRewardsInfoRecordExecutor{
    *postgresBigBangClient}
  purchaseMPsRecordExecutor := purchase_mps_record_config.PurchaseMPsRecordExecutor{
    *postgresBigBangClient}
  purchaseMPsRecordExecutor.UpsertPurchaseMPsRecordTx(purchaseMPsRecord)

  actorRewardsInfoRecordExecutor.AddActorMilestonePointsFromOthersTx(
    purchaseMPsRecord.Purchaser, feed_attributes.MilestonePoint(purchaseMPsRecord.MPs))

  postgresBigBangClient.Commit()
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
