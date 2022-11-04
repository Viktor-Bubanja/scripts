# scripts

The repo currently contains two different little scripts I use to automate parts of my workflow.

*Note: you need to have Go installed to run them.*

### Changing git branches
When this script is run, it prints out the last 10 branches I used in the terminal.
You can then enter a number between 0-9 to change to the specified branch (or the Enter key to exit).
This is useful for quickly navigating between several branches.

Here is the alias I added to my `~/.zshrc` file to run the script easily from anywhere:
```
alias lnb="go run ~/Documents/scripts/change_git_branch/change_git_branch.go"
```

### Copying the last commit hash
When addressing PR feedback, it can be useful to copy paste the commit hash that implements the author's suggestion to streamline the review process.
When this script is run, it simply copies the commit hash of the last commit.

The alias I use:
```
alias cph="go run ~/Documents/scripts/copy_commit_hash/copy_commit_hash.go"
```
