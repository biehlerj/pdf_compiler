package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}
	compileDir := dirname + "/dev/git.sr.ht/biehlerj/Awesome-CV"

	chDirErr := os.Chdir(compileDir)

	if chDirErr != nil {
		log.Fatal(chDirErr)
	}
	fmt.Println(os.Getwd())
	cmd := exec.Command("make", "all")
	cmd.Dir = compileDir
	var out bytes.Buffer
	cmd.Stdout = &out
	cmdErr := cmd.Run()
	if cmdErr != nil {
		log.Fatal(cmdErr)
	}
	fmt.Println(out.String())
}
