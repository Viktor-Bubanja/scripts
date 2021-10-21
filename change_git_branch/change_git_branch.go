package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

const copy_character = "c"

func main() {
	c := color.New(color.FgGreen)
	num_branches := 10
	command := fmt.Sprintf("git for-each-ref --sort='-authordate:iso8601' refs/heads | head -n%d", num_branches)
	stdout, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	branches := strings.Split(string(stdout), "\n")
	num_branches = cap(branches) - 1
	if num_branches <= 0 {
		return
	}
	formatted_branches := [10]string{}
	for i, branch := range branches[:num_branches] {
		branch_name := strings.Fields(branch)[2]
		branch_prefix := "refs/heads/"
		formatted_branch := strings.TrimPrefix(branch_name, branch_prefix)
		formatted_branches[i] = formatted_branch
		c.Print("[" + fmt.Sprint(i) + "] ")
		fmt.Println(formatted_branch)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Change branch?")
	input, _ := reader.ReadString('\n')
	if string(input[0]) == copy_character {
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
