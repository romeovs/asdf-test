package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hello from go!")

	cmd := exec.Command("node", "./test.js")
	cmd.Env = []string{
		"PATH=" + os.Getenv("PATH"),

		// if you uncomment the following line it works:
		// "ASDF_DATA_DIR=" + os.Getenv("ASDF_DATA_DIR"),
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
