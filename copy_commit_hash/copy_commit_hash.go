package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	stdout, err := exec.Command("bash", "-c", "git log").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	log_lines := strings.Split(string(stdout), "\n")
	first_line := log_lines[0]
	message := strings.TrimSpace(log_lines[4])

	commit_hash := strings.Split(first_line, " ")[1]
	copy_command := fmt.Sprintf("echo -n %s | pbcopy", commit_hash)

	exec.Command("bash", "-c", copy_command).Output()
	fmt.Printf("Copied '%s', %s\n", message, commit_hash)
}
