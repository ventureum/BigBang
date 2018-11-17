package api

type URL string

const FeedSystemBaseURL URL = "https://7g1vjuevub.execute-api.ca-central-1.amazonaws.com"

var ProfileAlphaEndingPoint URL = BuildEndingPoint(FeedSystemBaseURL, AlphaStage, Profile)
var GetProfileAlphaEndingPoint URL = BuildEndingPoint(FeedSystemBaseURL, AlphaStage, GetProfile)
var GetBatchProfilesAlphaEndingPoint URL = BuildEndingPoint(FeedSystemBaseURL, AlphaStage, GetBatchProfiles)

var SetActorPrivateKeyAlphaEndingPoint URL = BuildEndingPoint(FeedSystemBaseURL, AlphaStage, SetActorPrivateKey)
var GetActorPrivateKeyAlphaEndingPoint URL = BuildEndingPoint(FeedSystemBaseURL, AlphaStage, GetActorPrivateKey)
var GetActorUuidFromPublicKeyAlphaEndingPoint URL = BuildEndingPoint(FeedSystemBaseURL, AlphaStage, GetActorUuidFromPublicKey)

var FeedPostAlphaEndingPoint URL = BuildEndingPoint(FeedSystemBaseURL, AlphaStage, FeedPost)
var GetFeedPostAlphaEndingPoint URL = BuildEndingPoint(FeedSystemBaseURL, AlphaStage, GetFeedPost)

const TCRBaseURL URL = "https://mfmybdhoea.execute-api.ca-central-1.amazonaws.com"


var NewProjectAlphaEndingPoint URL = BuildEndingPoint(TCRBaseURL, AlphaStage, NewProject)
var GetProjectAlphaEndingPoint URL = BuildEndingPoint(TCRBaseURL, AlphaStage, GetProject)

var NewProjectBetaEndingPoint URL = BuildEndingPoint(TCRBaseURL, BetaStage, NewProject)
var GetProjectBetaEndingPoint URL = BuildEndingPoint(TCRBaseURL, BetaStage, GetProject)