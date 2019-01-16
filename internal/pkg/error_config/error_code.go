package error_config

type ErrorCode string

const InvalidAuthAccess ErrorCode = "InvalidAuthAccess"
const InvalidAuthRegister ErrorCode = "InvalidAuthRegister"

// For feed
const GetStreamClientConnectionError ErrorCode = "GetStreamClientConnectionError"
const InsufficientWaitingTimeToRefuel ErrorCode = "InsufficientWaitingTimeToRefuel"
const InsufficientReputaionsAmount ErrorCode = "InsufficientReputaionsAmount"
const InsufficientFuelAmount ErrorCode = "InsufficientFuelAmount"
const NoActorExisting ErrorCode = "NoActorExisting"
const NoPrincipalIdExisting ErrorCode = "NoPrincipalIdExisting"
const InvalidPrincipalId ErrorCode = "InvalidPrincipalId"
const NoReDeemBlockInfoRecordExisting ErrorCode = "NoReDeemBlockInfoRecordExisting"
const NoActorExistingForPublicKey ErrorCode = "NoActorExistingForPublicKey"
const NoPrivateKeyExistingForActor ErrorCode = "NoPrivateKeyExistingForActor"

const NoPostHashExisting ErrorCode = "NoPostHashExisting"
const ExceedingUpvoteLimit ErrorCode = "ExceedingUpvoteLimit"
const ExceedingDownvoteLimit ErrorCode = "ExceedingDownvoteLimit"
const General ErrorCode = "General"
const EmptyPublicKey ErrorCode = "EmptyPublicKey"
const InvalidActorType ErrorCode = "InvalidActorType"
const InvalidPostType ErrorCode = "InvalidPostType"
const InvalidMilestonePoints ErrorCode = "InvalidMilestonePoints"
const WalletAddressAlreadyExisting ErrorCode = "WalletAddressAlreadyExisting"
const NoWalletAddressExisting ErrorCode = "NoWalletAddressExisting"

// For TCR
const NoProjectIdExisting ErrorCode = "NoProjectIdExisting"
const NoObjectiveIdExisting ErrorCode = "NoObjectiveIdExisting"
const NoMilestoneIdExisting ErrorCode = "NoMilestoneIdExisting"
const NoRatingVoteVoterExisting ErrorCode = "NoRatingVoteVoterExisting"
const NoProxyUUIDExisting ErrorCode = "NoProxyUUIDExisting"
const ProxyUUIDAlreadyExisting ErrorCode = "ProxyUUIDAlreadyExisting"
const TotalProxyVotingPercentageExceeding100 ErrorCode = "TotalProxyVotingPercentageExceeding100"
const RatingVoteExceedingLimitedVotingTimes ErrorCode = "RatingVoteExceedingLimitedVotingTimes"
const MilestoneInvalidForUpdating ErrorCode = "MilestoneInvalidForUpdating"
const MilestoneValidatorAlreadyExisting ErrorCode = "MilestoneValidatorAlreadyExisting"
const ProjectAdminReassign ErrorCode = "ProjectAdminReassign"
