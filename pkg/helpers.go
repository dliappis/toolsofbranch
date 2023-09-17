package pkg

import (
	"fmt"
	"log"
	"math"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	gosemver "golang.org/x/mod/semver"
)

var maxVer string = fmt.Sprintf("%d.%d.%d", math.MaxInt32, math.MaxInt32, math.MaxInt32)

// GitGitBranch returns the current (short) branch name under the specified curDir. Equivalent to "cd curDir && git branch"
func GetGitBranch(curDir string) string {
	r, err := git.PlainOpen((curDir))
	if err != nil {
		log.Fatalf("Unable to read [%s] as a git repository.\nError: %s\n", curDir, err)
	}

	var p plumbing.ReferenceName = "HEAD"
	var pRef *plumbing.Reference
	pRef, err = r.Reference(p, true)
	if err != nil {
		log.Fatalf("Unable to obtain current branch in git checkout [%s].\nError: %s\n", curDir, err)
	}

	return pRef.Name().Short()
}

// EffectiveBranch validates that a branch name is valid semver and returns it unmodified.
// If not, it returns what is effectively the largest semver possible.
func EffectiveBranch(gitBranch string) string {
	effectiveBranch := gitBranch

	if !gosemver.IsValid("v" + gitBranch) {
		effectiveBranch = maxVer
	}

	return effectiveBranch
}
