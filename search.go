package main

import (
  "github.com/olivere/elastic"
  "fmt"
)

func doc() string {
  client, err := elastic.NewClient()
  if err != nil {
    // Handle error
    panic(err)
  }

  get1, err := client.
                Get().
                Index("destination-translations").
                Type("translation").
                Id("6734").
                Do()

  if err != nil {
      // Handle error
      panic(err)
  }
  if get1.Found {
    return fmt.Sprintf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
  }
  return ""
}