// Author: beaker
// This script is designed to load test web servers
// Send numerous concurrent requests and see which 
// respond back with OK and monitor when they start to fail

package main

import (
  "fmt"
  "log"
  "os"
  "io"
  "flag"
  "net/http"
)

var qty int
var url string
var logging bool

type Stats struct {
  Success int
  Failure int
}

type Result struct {
  Status int // 0 for failure, 1 for success
  Data *http.Response
}

func init() {
  flag.StringVar(&url, "url", "http://localhost", "the url to ddos")
  flag.IntVar(&qty, "qty", 10, "how many requests to make")
  flag.BoolVar(&logging, "log", false, "set logging on/off")
  banner := `
                 ,mmm^>
   \|  ddos  __   W-W'  ___
  `
  fmt.Println(banner)
}

func main() {
  // Setup command line arguments
  flag.Parse()

  summary := Stats{
    Success:0,
    Failure: 0,
  }

  // Setup logging
  if logging {
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

    if err != nil {
      log.Fatal(err)
    }
    log.SetOutput(file)
  } else {
    log.SetOutput(io.Discard)
  }

  // Setup channels for tests and results
  tests := make(chan int)
  results := make(chan Result)

  // Create a go routine for each request
  for i:= 0; i < qty; i++ {
    go makeRequest(url, tests, results)
  }

  // Make the requests
  go func() {
    for i:=0; i < qty; i++ {
      tests <- i
    }
  }()

  //TODO make this a switch statement and remove the n.Status == 1 ?
  // probz clean this up when handle various status codes
  // and more detailed stats
  for n := range results {
    if n.Status == 1 {
      summary.Success += 1
    } else {
      summary.Failure += 1
    }
    total := summary.Success + summary.Failure
    fmt.Printf("\rRequests Made: %d/%d", total, qty)
    if total == qty {
      close(results)
    }
  }
  fmt.Println("\n")
  fmt.Printf("Success: %d\n", summary.Success)
  fmt.Printf("Failure: %d", summary.Failure)
}

func makeRequest(url string, tests chan int, results chan Result) {
    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
      fmt.Println(err)
    }

    client := &http.Client{}

    res, err := client.Do(req)
    // TODO convert to switch statement and handle different error codes?
    if err != nil {
      // If error due to tcp reset, etc.
      log.Println(err.Error())
      results <- Result{Status: 0, Data: &http.Response{}} 
    } else {
      log.Println(res)
      results <- Result{Status: 1, Data: res}
      defer res.Body.Close()
    }
}
