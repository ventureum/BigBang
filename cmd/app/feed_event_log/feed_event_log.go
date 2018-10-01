package main

import (
  "os"
  "log"
  "BigBang/internal/platform/postgres_config/feed/client_config"
  "BigBang/internal/platform/eth_config"
  "BigBang/internal/platform/getstream_config"
)

func main() {
  forumAddress := os.Getenv("FORUM_ADDRESS")

  if forumAddress == "" {
   log.Fatal("forum address is not set yet")
  }

  log.Printf("Get Forum Address: %s\n", forumAddress)

  log.Println("Connecting to Ethereum EthClient")
  ethClient := eth_config.ConnectEthClient()

  log.Println("Connecting to GetStream Client")
  getStreamClient := getstream_config.ConnectGetStreamClient()

  log.Println("Connecting to Postgres Client")
  postgresClient := client_config.ConnectPostgresClient()

  log.Printf("Subscribing to logs at Forum Address: %s\n", forumAddress)
  ethClient.SubscribeFilterLogs(forumAddress, getStreamClient, postgresClient)
}
