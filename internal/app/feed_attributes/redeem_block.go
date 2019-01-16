package feed_attributes

import (
	"os"
	"strconv"
	"time"
)

type RedeemBlock int64

const DefaultRedeemBlockLength int64 = 60 * 60 * 24 * 7

var RedeemBlockLength = LoadRedeemBlockLengthEnv()

func CreateRedeemBlockFromUnix(unix int64) RedeemBlock {
	return RedeemBlock(unix / RedeemBlockLength)
}

func MoveToNextNRedeemBlock(n int64) RedeemBlock {
	return RedeemBlock(time.Now().UTC().Unix()/RedeemBlockLength + n)
}

func (redeemBlock RedeemBlock) ConvertToTime() time.Time {
	convertedTime := time.Unix(int64(redeemBlock)*RedeemBlockLength, 0)
	return convertedTime.In(time.UTC)
}

func LoadRedeemBlockLengthEnv() int64 {
	var redeemBlockLength int64 = DefaultRedeemBlockLength
	val := os.Getenv("REDEEM_BLOCK_LENGTH")
	if val != "" {
		redeemBlockLength, _ = strconv.ParseInt(val, 10, 64)
	}
	return redeemBlockLength
}
