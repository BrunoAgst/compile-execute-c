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

	cmdNasm := exec.Command("/opt/homebrew/bin/nasm", asm, "-f", "bin", "-o", bin)
	out, err := cmdNasm.CombinedOutput()

	if err != nil {
		log.Printf("%s\n", out)
		return
	}

	qemuCommand := fmt.Sprintf("file=%s,format=raw,index=0,if=floppy", bin)

	cmdQemu := exec.Command("/opt/homebrew/bin/qemu-system-i386", "-drive", qemuCommand)

	cmdQemu.Start()
}
