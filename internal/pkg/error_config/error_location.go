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
const ActorTypeLocation ErrorLocation = "ActorTypeLocation"
const PostTypeLocation ErrorLocation = "PostTypeLocation"

// For TCR
const ProjectRecordLocation ErrorLocation = "ProjectRecordLocation"
const MilestoneRecordLocation ErrorLocation = "MilestoneRecordLocation"
const MilestoneValidatorRecordLocation ErrorLocation = "MilestoneValidatorRecordLocation"
const ObjectiveRecordLocation ErrorLocation = "ObjectiveRecordLocation"
const ProxyRecordLocation ErrorLocation = "ProxyRecordLocation"
const RatingVoteRecordLocation ErrorLocation = "RatingVoteRecordLocation"
const ActorDelegateVotesAccountRecordLocation ErrorLocation = "ActorDelegateVotesAccountRecordLocation"
const PrincipalProxyVotesRecordLocation ErrorLocation = "PrincipalProxyVotesRecordLocation"
