package exiftool

import (
	"bytes"
	"log"
	"os/exec"
)

const name = "exiftool"

func do(arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	var output bytes.Buffer
	cmd.Stdout = &output
	if err := cmd.Start(); err != nil {
		log.Fatalln("Error: ", cmd.String(), err)
	}
	cmd.Wait()
	log.Println("OK ", cmd.String())
	return output.Bytes(), nil
}
