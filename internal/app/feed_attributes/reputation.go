package feed_attributes

import (
  "math/big"
  "strconv"
)


type Reputation int64


const REPUTATION_BASE = 100

var PostReputationCost Reputation = 10 * REPUTATION_BASE
var ReplyReputationCost Reputation = 1 * REPUTATION_BASE
var AduitReputationCost Reputation = 100 * REPUTATION_BASE


func PenaltyForPostType(postType PostType, counter int64) Reputation {
  var reputations Reputation
  switch postType {
    case PostPostType:
        reputations = PostReputationCost.MulByPower(big.NewInt(2), big.NewInt(counter))
    case ReplyPostType:
        reputations = ReplyReputationCost.MulByPower(big.NewInt(2), big.NewInt(counter))
    case AuditPostType:
        reputations = AduitReputationCost.MulByPower(big.NewInt(2), big.NewInt(counter))
  }
  return reputations
}

func CreateReputationFromStr(rep string) Reputation {
  i, _ := strconv.ParseInt(rep, 10, 64)
  return Reputation(i)
}

func PenaltyForVote(base Reputation, counter int64) Reputation {
  return base.MulByPower(big.NewInt(2), big.NewInt(counter))
}

func BigIntToReputation(num *big.Int) Reputation {
  return Reputation(num.Int64())
}

func (reputation Reputation) Value() int64 {
  return int64(reputation)
}

func (reputation Reputation) ToBigInt() *big.Int {
  return big.NewInt(reputation.Value())
}

func (reputation Reputation) AddToReputations(reputationToAdd Reputation) Reputation {
  num := new(big.Int)
  num.Add(reputation.ToBigInt(), reputationToAdd.ToBigInt())
  return BigIntToReputation(num)
}

func (reputation Reputation) SubReputations(reputationToSub Reputation) Reputation {
  num := new(big.Int)
  num.Sub(reputation.ToBigInt(), reputationToSub.ToBigInt())
  return BigIntToReputation(num)
}

func (reputation Reputation) MulByPower(base *big.Int, factor *big.Int) Reputation {
  num := new(big.Int)
  numInt := new(big.Int).Exp(base, factor, nil )
  num.Mul(reputation.ToBigInt(), numInt)

  return BigIntToReputation(num)
}

func (reputation Reputation) Sign() int {
  num := reputation.ToBigInt()
  return num.Sign()
}

func (reputation Reputation) Abs() int64 {
  num := new(big.Int)
  return num.Abs(reputation.ToBigInt()).Int64()
}

func (reputation Reputation) Neg() Reputation {
  num := new(big.Int)
  return BigIntToReputation(num.Neg(reputation.ToBigInt()))
}
