package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const copyCharacter = "c"
const colorReset = "\033[0m"
const colorGreen = "\033[32m"

func main() {
	numBranches := 10
	command := fmt.Sprintf("git for-each-ref --sort='-authordate:iso8601' refs/heads | head -n%d", numBranches)
	stdout, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	branches := strings.Split(string(stdout), "\n")
	numBranches = cap(branches) - 1
	if numBranches <= 0 {
		return
	}
	formatted_branches := [10]string{}
	for i, branch := range branches[:numBranches] {
		branch_name := strings.Fields(branch)[2]
		branch_prefix := "refs/heads/"
		formatted_branch := strings.TrimPrefix(branch_name, branch_prefix)
		formatted_branches[i] = formatted_branch
		fmt.Print(string(colorGreen), "["+fmt.Sprint(i)+"] ")
		fmt.Println(string(colorReset), formatted_branch)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Change branch?")
	input, _ := reader.ReadString('\n')
	if string(input[0]) == copyCharacter {
		number, _ := strconv.ParseInt(string(input[1]), 10, 64)
		branch_to_copy := formatted_branches[number]
		copy_command := fmt.Sprintf("echo %s | pbcopy", branch_to_copy)
		exec.Command("bash", "-c", copy_command).Output()
	} else {
		number, _ := strconv.Atoi(string(input[0]))
		branch_to_checkout := formatted_branches[number]
		checkout_command := fmt.Sprintf("git checkout %s", branch_to_checkout)
		exec.Command("bash", "-c", checkout_command).Output()
	}
}
