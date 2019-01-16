package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"log"
	"strconv"
	"strings"
)

func urlSafe(src string) string {
	src = strings.Replace(src, "+", "-", -1)
	src = strings.Replace(src, "/", "_", -1)
	src = strings.Trim(src, "=")
	return src
}

func CryptoToken(id string, secret string) string {
	hash := sha1.New()
	hash.Write([]byte(secret))
	mac := hmac.New(sha1.New, hash.Sum(nil))
	mac.Write([]byte(id))
	digest := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return urlSafe(digest)
}

func Keccak256Hash(data ...[]byte) (h common.Hash) {
	d := sha3.NewKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	d.Sum(h[:0])
	return h
}

func Base64EncodeInt64(num int64) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(num, 10)))
	return encoded
}

func Base64DecodeToInt64(encoded string) int64 {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Panicf("decode error: %+v\n", err)
	}

	i, err := strconv.ParseInt(string(decoded), 10, 64)
	if err != nil {
		log.Panicf("parsing error: %+v\n", err)
	}
	return i
}

func Base64EncodeStr(str string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return encoded
}

func Base64DecodeToString(encoded string) string {
	if encoded == "" {
		return ""
	}
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Panicf("decode error: %+v\n", err)
	}
	return string(decoded)
}

func GenerateIdByInt64AndStr(firstField int64, secondField string) string {
	return strconv.FormatInt(firstField, 10) + ":" + secondField
}

func Base64EncodeIdByInt64AndStr(firstField int64, secondField string) string {
	return Base64EncodeStr(GenerateIdByInt64AndStr(firstField, secondField))
}
