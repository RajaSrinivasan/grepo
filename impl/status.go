package impl

import (
	"log"
	"os"
	"os/exec"

	"github.com/RajaSrinivasan/grepo/impl/repo"
)

func statusProjectGroup(prjg *repo.ProjectGroup) {
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	for _, prj := range prjg.Projects {
		log.Printf("Status in %s", prj.Path)
		os.Chdir(prj.Path)
		cmd := exec.Command("git", "status")
		result, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("%s", err)
		}
		log.Printf("%s", result)
	}
}

func Status(manifest *repo.Manifest) {
	if Verbose {
		log.Printf("Status")
	}

	statusProjectGroup(&manifest.Private)
}
