package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

func main(){
  cmd, err := exec.Command("arduino-cli", "board", "list").Output()
  if err != nil{
    log.Fatal("err", err.Error())
    return
  }

  match := regexp.MustCompile("a([a-z]+):([a-z]+):([a-z]+)")
  fqbn := match.FindString(string(cmd))
  fmt.Println(fqbn)
}

