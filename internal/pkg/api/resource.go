package api

type Resource string

// For Feed
const Profile Resource = "profile"
const GetProfile Resource = "get-profile"
const GetBatchProfiles Resource = "get-batch-profiles"

const SetActorPrivateKey Resource = "set-actor-private-key"
const GetActorPrivateKey Resource = "get-actor-private-key"
const GetActorUuidFromPublicKey Resource = "get-actor-uuid-from-public-key"

const FeedPost Resource = "feed-post"
const GetFeedPost Resource = "get-feed-post"
const GetBatchPosts Resource = "get-batch-posts"
const GetRecentPosts Resource = "get-recent-posts"

const AttachSession Resource = "attach-session"
const DeactivateActor Resource = "deactivate-actor"
const DevRefuel Resource = "dev-refuel"

const FeedUpvote Resource = "feed-upvote"

const GetRecentVotes Resource = "get-recent-votes"
const GetSession Resource = "get-session"

const Refuel Resource = "refuel"

// For TCR

const NewProject Resource = "new-project"
const GetProject Resource = "get-project"
