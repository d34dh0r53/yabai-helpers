package common

import (
	"fmt"
	"os/exec"
)

func dir_list() {
	cmd := exec.Command("ls", "-ltr")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
	} else {
		fmt.Println("Command successfully executed")
		fmt.Println("Output:", string(output))
	}
	return 0
}
