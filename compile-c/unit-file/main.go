package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main(){
  var fileName = os.Args[1]
  var fileNameReplace = strings.Replace(fileName, ".c", "", 1)
  
  //compile file
  compile := exec.Command("/usr/bin/gcc", fileName, "-o", fileNameReplace)
  _, err := compile.Output()
  if err != nil {
    fmt.Println("file not found", err.Error())
    return
  }

  //run file
  run := exec.Command(strings.Join([]string{"./", fileNameReplace}, ""))
  stout, err2 := run.Output()
  if err2 != nil {
    fmt.Println("compile error", err2.Error())
    return
  }
  fmt.Println(string(stout))
  
  //delete file
  if len(os.Args) > 2 && os.Args[2] == "-d" {
    return
  }
  
  err3 := os.Remove(fileNameReplace)
  if err3 != nil {
    fmt.Println("delete file error", err3.Error())
    return
  }
}    

