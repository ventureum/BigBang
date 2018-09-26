package api

type URL string

const FeedSystemBaseURL URL = "https://7g1vjuevub.execute-api.ca-central-1.amazonaws.com"

var ProfileAlphaEndingPoint URL = BuildEndingPoint(FeedSystemBaseURL, AlphaStage, Profile)
var GetProfileAlphaEndingPoint URL = BuildEndingPoint(FeedSystemBaseURL, AlphaStage, GetProfile)
