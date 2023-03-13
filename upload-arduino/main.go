package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main(){

  var projectName = os.Args[1]

  cmd, err := exec.Command("arduino-cli", "board", "list").Output()
  
  if err != nil{
    log.Fatal("error listBoard: ", err.Error())
    return
  }

  fqbn := getFqbn(cmd)
  
  compile(fqbn, projectName)

  upload(fqbn, cmd, projectName)
  
  log.Println("completed sucessfully")
}

func getFqbn(command []byte) string {
  match := regexp.MustCompile("a([a-z]+):([a-z]+):([a-z]+)")
  fqbn := match.FindString(string(command))
  
  return fqbn
}

func compile(fqbn string, name string) {
  cmd, err := exec.Command("arduino-cli", "compile", "--fqbn", fqbn, name).Output()

  if err != nil {
    log.Fatal("error Compile: ", err.Error())
  }

  log.Println(string(cmd))
}

func upload(fqbn string, cmd []byte, name string) {
  match := regexp.MustCompile("/dev/cu.([a-z0-9]+)")
  pathBoard := match.FindString(string(cmd))

  _, err := exec.Command("arduino-cli", "upload", "-p", pathBoard, "--fqbn", fqbn, name).Output()

  if err != nil {
    log.Fatal("err upload: ", err.Error())
  }
}
