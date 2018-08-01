package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

var (
	version   = "1.0.0"
	buildDate = ""
	commit    = ""
)

func versionText() string {
	msg := fmt.Sprintf("%-12s %s", "Version:", version)
	if buildDate != "" {
		msg += fmt.Sprintf("\n%-12s %s", "Built on:", buildDate)
	}
	if commit != "" {
		msg += fmt.Sprintf("\n%-12s %s", "Git Commit:", commit)
	}
	return msg
}

func main() {
	cmd := NewBcryptGenerateCmd()
	parser := flags.NewParser(cmd, flags.HelpFlag|flags.PassDoubleDash)
	parser.LongDescription = versionText()
	args, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := cmd.Execute(args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
