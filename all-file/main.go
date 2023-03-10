package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main(){
  
  files := readFiles()

  filesCompiled := compileFiles(files)
  
  compileFinalFile(filesCompiled)

  removeFiles(filesCompiled)

}

func readFiles() []string {
  files, err := ioutil.ReadDir(".")
  
  if err != nil {
    log.Fatal("ReadDir err :::", err.Error())
  }
  filesName := []string{}

  for _, file := range files{
    if(strings.Contains(file.Name(), ".c")){
      filesName = append(filesName, file.Name())
    }
  }
  return filesName
}

func compileFiles(files []string) []string {
  filesCompiled := []string{}

  for _, file := range files{
    var fileOutput = strings.Replace(file, ".c", ".o", 1)
  
    _, err := exec.Command("/usr/bin/gcc", "-c", file, "-o", fileOutput).Output()
    
    if err != nil{
      log.Fatal("Compile err :::", err.Error())
    }
    
    filesCompiled = append(filesCompiled, fileOutput)
  }

  return filesCompiled
}

func compileFinalFile(files []string) {
  files = append(files, "-o")
  files = append(files, "output")

  _, err := exec.Command("/usr/bin/gcc", files...).Output()

  if err != nil{
    log.Fatal("Compile final err:::", err.Error())
  }
}

func removeFiles(files []string) {
  for _, file := range files{
  
    err := os.Remove(file)
    
    if err != nil {
      log.Fatal("Remove file err:::", err.Error())
    }
  }
}
