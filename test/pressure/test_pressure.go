package main

import (
  "BigBang/internal/pkg/api"
  "fmt"
  "strconv"
  "sync"
  "log"
  "math/rand"
)

const TestSize = 500

func RunJob(wg *sync.WaitGroup, index int) {
  defer wg.Done()
  CreateProfile(index)
  DevRefuel(index)
  UpsertPost(index)
  FeedUpvote(index)
  FeedUpvote(index)
}

func CreateProfile(index int) {
    actor := "testActor" +  strconv.Itoa(index)
    message := api.Message(map[string]interface{}{
      "actor": actor,
      "userType": "USER",
    })
    profileURL := api.BuildEndingPoint(api.FeedSystemBaseURL, api.AlphaStage, api.Profile)
    response := api.SendPost(message, profileURL)
    log.Printf("Response for actor %s: %+v\n",  actor, response)
}


func DevRefuel(index int) {
  actor := "testActor" +  strconv.Itoa(index)
  message := api.Message(map[string]interface{}{
    "actor": actor,
    "fuel": 10000,
    "reputation": 10000,
    "milestonePoints": 10000,
  })
  devRefuelURL := api.BuildEndingPoint(api.FeedSystemBaseURL, api.AlphaStage, api.DevRefuel)
  response := api.SendPost(message, devRefuelURL)
  log.Printf("Response for DevRefuel by actor %s: %+v\n",  actor, response)
}


func UpsertPost(index int) {
  actor := "testActor" +  strconv.Itoa(index)
  postHash := "testHash" + strconv.Itoa(index)
  message := api.Message(map[string]interface{}{
      "actor": actor,
      "boardId": "botTest2",
      "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
      "postHash":  postHash,
      "typeHash": "0x2fca5a5e",
      "content": map[string]interface{} {
        "title": "titleSample111",
        "text": "Introduction\nAs part of I'm he's![image](https://images.pexels.com/photos/991438/pexels-photo-991438.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260)“best practices” series. we wanted to describe our best practices for setting up! feeds for a simple photo sharing application. These feeds allow users to upload photos, follow other users, like photos, and see notification and aggregated feeds.\n\nOf course, how you use Stream depends greatly on the kind of application you’re building and your intention for the data, so the best practices we define for this photo sharing application may not line up exactly with your project or business model. Feel free to reach out to our support team if you have questions.\n\nRegardless of your use case, our best practice when accessing user feeds is to utilize a UUID for identifying things in your app (users, photos, etc) to avoid collisions. For the sake of this blog post, the examples below will use usernames to make ![image](https://images.pexels.com/photos/991438/pexels-photo-991438.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260) it easier to follow along.\n\nWe created several feeds for our photo sharing app: two flat feeds, one aggregated feed and one notification feed.",
        "image": "https://images.pexels.com/photos/991438/pexels-photo-991438.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260",
        "subtitle": "Introduction As part of I'm he's “best practices” series. we wanted to describe our best practices for setting up! feeds for a simple photo sharing application. These feeds allow ",
      },
  })
  feedPostURL := api.BuildEndingPoint(api.FeedSystemBaseURL, api.AlphaStage, api.FeedPost)
  response := api.SendPost(message, feedPostURL)
  log.Printf("Response for Creating Post  %s by actor %s: %+v\n",  postHash, actor, response)
}

func FeedUpvote(index int) {
  actor := "testActor" +  strconv.Itoa(index)
  postHash := "testHash" + strconv.Itoa(rand.Intn(TestSize))
  value := 1 - rand.Intn(3)
  message := api.Message(map[string]interface{}{
    "actor": actor,
    "postHash": postHash,
    "value": value,
  })
  feedUpvoteURL := api.BuildEndingPoint(api.FeedSystemBaseURL, api.AlphaStage, api.FeedUpvote)
  response := api.SendPost(message, feedUpvoteURL)
  log.Printf("Response for Upvoting Post %s by actor %s with value %d: %+v\n",  postHash, actor, value, response)
}

func main() {
  var wg sync.WaitGroup
  for i := 0; i < TestSize; i++ {
      log.Printf(" starting Job %d \n",  i)
      wg.Add(1)
      go RunJob(&wg, i)
  }
  log.Println("Main: waiting for workers to finish")
  wg.Wait()
  fmt.Println("Main: completed")
}
