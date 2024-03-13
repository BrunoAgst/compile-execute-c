package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var fileName = os.Args[1]

	asm := fileName + ".asm"
	bin := fileName + ".bin"

	_, err := exec.Command("/opt/homebrew/bin/nasm", asm, "-f", "bin", "-o", bin).Output()

	if err != nil {
		log.Fatal("Error nasm: ", err.Error())
		return
	}

	qemuCommand := fmt.Sprintf("file=%s,format=raw,index=0,if=floppy", bin)

	cmd := exec.Command("/opt/homebrew/bin/qemu-system-i386", "-drive", qemuCommand)

	cmd.Start()
}
