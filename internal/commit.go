package internal

import (
	"fmt"
	"os/exec"
	"strings"
)

type CommitObject struct {
	commitMessage   string
	commitType      string
	breakingChanges bool
	hasScope        bool
	scopeMessage    string
}

func NewCommitMessage(commitType string, commitMessage string) *CommitObject{
	commit := CommitObject{}
	commit.commitType = commitType
	commit.commitMessage = commitMessage
	commit.breakingChanges = false
	commit.hasScope = false
	commit.scopeMessage = ""
	return &commit
}

func (c *CommitObject) HasBreakingChanges(){
	c.breakingChanges = true
}

func (c *CommitObject) AddScope(scopeMessage string){
	c.hasScope = true
	c.scopeMessage = scopeMessage
}

func Commit(c *CommitObject) {
	/* Takes a commit message (like the one you pass to git) and a
	conventional commit type. For example fix, feat or chore
	*/
	var sb strings.Builder
	sb.WriteString(c.commitType)
	if c.hasScope {
		sb.WriteString("(")
		sb.WriteString(c.scopeMessage)
		sb.WriteString(")")
	}
	if c.breakingChanges {
		sb.WriteString("!")
	}
	sb.WriteString(": ")
	sb.WriteString(c.commitMessage)

	execCmd := exec.Command("git", "commit", "-m", sb.String())
	stdout, err := execCmd.Output()

	if err != nil {
		fmt.Println("failed to commit double check you have staged your changes")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
}
