package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/sagarsearce/gospin"
	"github.com/spf13/cobra"
)

var (
	remoteUrl   string
	branch      string
	commitMsg   string
	cm          string
	rootCommand = cobra.Command{
		Use:   "it",
		Short: "speed up git command.",
		Long: `--it-- is tool that will speed up you git commad like there ae
		use case when you initialize your repo you will run 4 git command in --it-- you just
		have to run one command with flag    .`,
	}
	ft = cobra.Command{
		Use:   "ft",
		Short: "ft stand for fitst time, when you set local repo first time",
		Long:  `ft will initialize a repo, stage you code that set up remote and push code there in master branch`,
		Run: func(cmd *cobra.Command, args []string) {
			message := "initlizing git repo ..."
			spinner := gospin.CreateLoder("50ms", `-\|/`)
			go spinner.StartLoading(message)
			in := []string{"git", "init"}
			executeCommand(in)

			message = "adding remote.."
			addR := []string{"git", "remote", "add", "origin", remoteUrl}
			executeCommand(addR)

			message = "git add ."
			add := []string{"git", "add", "."}
			executeCommand(add)

			message = "git commit.."
			c := []string{"git", "commit", "-m", commitMsg}
			executeCommand(c)

			message = "pushing..."
			p := []string{"git", "push", "-u", "origin", "master"}
			executeCommand(p)

			message = "logging.."
			l := []string{"git", "log"}
			executeCommand(l)

			spinner.StopLoading()
		},
	}
	clean = cobra.Command{
		Use:   "clean",
		Short: "add,commit and push ",
		Long:  `clean will add commit and push your code if push give error that means conflicts are there.`,
		Run: func(cmd *cobra.Command, args []string) {
			message := "adding.."
			spinner := gospin.CreateLoder("50ms", `-\|/`)
			go spinner.StartLoading(message)

			message = "git add ."
			add := []string{"git", "add", "."}
			executeCommand(add)

			message = "git commit.."
			c := []string{"git", "commit", "-m", cm}
			executeCommand(c)

			message = "pushing..."
			p := []string{"git", "push", "-u", "origin", branch}
			executeCommand(p)

			spinner.StopLoading()
		},
	}
)

func init() {
	rootCommand.AddCommand(&ft)
	rootCommand.AddCommand(&clean)

	ft.Flags().StringVarP(&remoteUrl, "remote", "r", "", "add remote url for the repo")
	ft.Flags().StringVarP(&commitMsg, "commitmsg", "m", "first commit", "add commit message")
	ft.MarkFlagRequired("remote")
	clean.Flags().StringVarP(&branch, "branch", "b", "master", "initial branch name")
	clean.Flags().StringVarP(&cm, "commitmsg", "m", "commit message", "add commit message")
	clean.MarkFlagRequired("commitmsg")
	clean.MarkFlagRequired("branch")
}

func executeCommand(p []string) {
	terminal := exec.Command(p[0], p[1:]...)
	out, err := terminal.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
