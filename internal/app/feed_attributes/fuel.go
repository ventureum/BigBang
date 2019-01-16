package feed_attributes

import (
	"math/big"
	"os"
	"strconv"
)

type Fuel int64

var MuMaxFuel Fuel = loadEnv("MuMaxFuel")
var MuMinFuel Fuel = loadEnv("MuMinFuel")
var PostFuelCost Fuel = loadEnv("PostFuelCost")
var ReplyFuelCost Fuel = loadEnv("ReplyFuelCost")
var AuditFuelCost Fuel = loadEnv("AuditFuelCost")
var BetaMax Fuel = loadEnv("BetaMax")
var MaxFuelForFuelUpdateInterval Fuel = loadEnv("MAX_FUEL_FOR_FUEL_UPDATE_INTERVAL")

func loadEnv(key string) Fuel {
	val, _ := strconv.Atoi(os.Getenv(key))
	return Fuel(val)
}

func FuelsPenaltyForPostType(postType PostType, counter int64) Fuel {
	var fuels Fuel
	switch postType {
	case PostPostType:
		fuels = PostFuelCost.MulByPower(big.NewInt(2), big.NewInt(counter))
	case ReplyPostType:
		fuels = ReplyFuelCost.MulByPower(big.NewInt(2), big.NewInt(counter))
	case AuditPostType:
		fuels = AuditFuelCost.MulByPower(big.NewInt(2), big.NewInt(counter))
	}
	return fuels
}

func BigIntToFuel(num *big.Int) Fuel {
	return Fuel(num.Int64())
}

func (fuel Fuel) Value() int64 {
	return int64(fuel)
}

func (fuel Fuel) ToBigInt() *big.Int {
	return big.NewInt(fuel.Value())
}

func (fuel Fuel) AddToFuels(fuelToAdd Fuel) Fuel {
	num := new(big.Int)
	num.Add(fuel.ToBigInt(), fuelToAdd.ToBigInt())
	return BigIntToFuel(num)
}

func (fuel Fuel) SubFuels(fuelToSub Fuel) Fuel {
	num := new(big.Int)
	num.Sub(fuel.ToBigInt(), fuelToSub.ToBigInt())
	return BigIntToFuel(num)
}

func (fuel Fuel) MulByPower(base *big.Int, factor *big.Int) Fuel {
	num := new(big.Int)
	numInt := new(big.Int).Exp(base, factor, nil)
	num.Mul(fuel.ToBigInt(), numInt)

	return BigIntToFuel(num)
}

func (fuel Fuel) Sign() int {
	num := fuel.ToBigInt()
	return num.Sign()
}

func (fuel Fuel) Abs() int64 {
	num := new(big.Int)
	return num.Abs(fuel.ToBigInt()).Int64()
}

func (fuel Fuel) Neg() Fuel {
	num := new(big.Int)
	return BigIntToFuel(num.Neg(fuel.ToBigInt()))
}
