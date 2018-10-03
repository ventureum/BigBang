package api

type Resource string

// For Feed
const AttachSession Resource = "attach-session"
const DeactivateActor Resource = "deactivate-actor"
const DevRefuel Resource = "dev-refuel"
const FeedPost Resource = "feed-post"
const FeedUpvote Resource = "feed-upvote"
const GetBatchPosts Resource = "get-batch-posts"
const GetFeedPost Resource = "get-feed-post"
const GetProfile Resource = "get-profile"
const GetRecentPosts Resource = "get-recent-posts"
const GetRecentVotes Resource = "get-recent-votes"
const GetSession Resource = "get-session"
const Profile Resource = "profile"
const Refuel Resource = "refuel"

// For TCR

const NewProject Resource = "new-project"
const GetProject Resource = "get-project"