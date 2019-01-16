package feed_attributes

import "math"

type Reputation int64

func (reputation Reputation) Value() int64 {
	return int64(reputation)
}

func (reputation Reputation) AwardFuels() Fuel {
	return Fuel(math.Max(float64(reputation), float64(MuMaxFuel)))
}
