package utils

import (
  "strings"
  "crypto/sha1"
  "crypto/hmac"
  "encoding/base64"
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
