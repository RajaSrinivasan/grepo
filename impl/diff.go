package impl

import (
	"log"
	"os"
	"os/exec"

	"github.com/RajaSrinivasan/grepo/impl/repo"
)

func diffProjectGroup(prjg *repo.ProjectGroup) {
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	for _, prj := range prjg.Projects {
		log.Printf("Diff in %s", prj.Path)
		os.Chdir(prj.Path)
		cmd := exec.Command("git", "diff", prj.Reference)
		result, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("%s", err)
		}
		log.Printf("%s", result)
	}
}

func Diff(manifest *repo.Manifest) {
	if Verbose {
		log.Printf("Diff")
	}

	diffProjectGroup(&manifest.Private)
}
