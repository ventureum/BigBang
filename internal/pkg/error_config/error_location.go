package error_config

type ErrorLocation string

// For feed
const ReputationsAccountLocation ErrorLocation = "ReputationsAccount"
const ActorRewardsInfoRecordLocation ErrorLocation = "ActorRewardsInfoRecordLocation"
const ProfileAccountLocation ErrorLocation = "ProfileAccount"
const PostRecordLocation ErrorLocation = "PostRecordLocation"
const PostRepliesRecordLocation ErrorLocation = "PostRepliesRecordLocation"
const ActorVotesCountersRecordLocation ErrorLocation = "ActorVotesCountersRecordLocation"
const PostVotesRecordLocation ErrorLocation = "PostVotesRecordLocation"
const SessionRecordLocation ErrorLocation = "SessionRecordLocation"
const RefuelRecordLocation ErrorLocation = "RefuelRecordLocation"

// For TCR
const ProjectRecordLocation ErrorLocation = "ProjectRecordLocation"