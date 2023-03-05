package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main(){
  var fileName = os.Args[1]
  
  //compile file
  compile := exec.Command("/usr/bin/gcc", fileName, "-o", "output")
  _, err := compile.Output()
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  //run file
  run := exec.Command("./output")
  stout, err2 := run.Output()
  if err2 != nil {
    fmt.Println(err2.Error())
    return
  }
  fmt.Println(string(stout))
  
  //delete file
  if len(os.Args) > 2 && os.Args[2] == "-d" {
    return
  }

  err3 := os.Remove("output")
  if err3 != nil {
    fmt.Println(err3.Error())
     return
  }
}    

