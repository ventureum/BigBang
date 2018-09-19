package main

import (
  "BigBang/internal/pkg/api"
  "fmt"
  "strconv"
  "sync"
  "log"
)

func profiles(wg *sync.WaitGroup, index int) {
    defer wg.Done()
    actor := "0xtest" +  strconv.Itoa(index)
    message := api.Message(map[string]interface{}{
      "actor": actor,
      "userType": "USER",
    })
    profileURL := api.BuildEndingPoint(api.FeedSystemBaseURL, api.AlphaStage, api.Profile)
    response := api.SendPost(message, profileURL)
    log.Printf("Response for actor %s: %+v\n",  actor, response)
}


func main() {
  var wg sync.WaitGroup
  numActors := 1000
  for i := 0; i < numActors; i++ {
      log.Printf(" starting request for actor index %d \n",  i)
      wg.Add(1)
      go profiles(&wg, i)
  }
  log.Println("Main: waiting for workers to finish")
  wg.Wait()
  fmt.Println("Main: completed")
}
