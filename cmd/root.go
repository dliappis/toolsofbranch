package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Masterminds/semver/v3"
	"github.com/dliappis/toolsofbranch/pkg"
	"github.com/spf13/cobra"
)

var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "toolsofbranch",
	Short: "A tool to determine applicable versions of external tools based on branch version",
	Long:  `A tool to help filtering a list of tools with semver constraints based on the branch of a checked out repository.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		gitBranch := pkg.GetGitBranch(args[1])
		version(args[0], gitBranch)
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func version(toolsfile string, gitBranch string) {
	var compatTools []string

	var d pkg.Dependencies
	toolsSpec, err := os.ReadFile(toolsfile)
	if err != nil {
		log.Fatal(err)
	}

	pkg.ReadTools(toolsSpec, &d)

	effectiveBranch := pkg.EffectiveBranch(gitBranch)
	curBranchVer, err := semver.NewVersion(effectiveBranch)
	if err != nil {
		log.Fatalf("The current branch name [%s] in git isn't compatible with semver!\nError: %s\n", curBranchVer, err)
	}

	for toolFamily, toolData := range d.Data {
		for toolName, toolConstraints := range toolData {
			c, err := semver.NewConstraint(toolConstraints)
			if err != nil {
				log.Fatalf("Unable to parse semver for %s/%s: %s.\nError: %s\n", toolFamily, toolName, toolConstraints, err)
			}
			if c.Check(curBranchVer) {
				compatTools = append(compatTools, toolName)
			}

		}
	}

	if Verbose {
		fmt.Printf("Report of eligible tools based on current branch [%s]\n", gitBranch)
		fmt.Println("=========================================================")
	}
	for _, tool := range compatTools {
		fmt.Println(tool)
	}
}
