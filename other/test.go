package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func main() {
  content, err := ioutil.ReadFile("filename")
  if err != nil {
      //Do something
  }
  lines := strings.Split(string(content), "\n")
  for _, line := range lines {
    fmt.Println(line)
  }
}
