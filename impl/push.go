package impl

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/RajaSrinivasan/grepo/impl/repo"
)

func pushProjectGroup(prjg *repo.ProjectGroup, message string, force bool) {
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	for _, prj := range prjg.Projects {
		log.Printf("Push in %s", prj.Path)
		os.Chdir(prj.Path)

		cmdstat := exec.Command("git", "status", "--porcelain")
		result, err := cmdstat.CombinedOutput()
		if err != nil {
			log.Printf("%s", err)
			continue
		}
		if len(result) < 1 {
			log.Printf("No changes to commit. Skipping")
			if force {
				log.Printf("Will commit anyway.")
			} else {
				continue
			}
		}
		log.Printf("%s", string(result))
		msgclean := fmt.Sprintf("\"%s\"", message)
		cmd := exec.Command("git", "commit", "-m", msgclean)

		result, err = cmd.CombinedOutput()
		if err != nil {
			log.Printf("%s", err)
			continue
		}
		log.Printf("%s", result)

		pushcmd := exec.Command("git", "push")
		result, err = pushcmd.CombinedOutput()
		if err != nil {
			log.Printf("%s", err)
		}
		log.Printf("%s", result)
	}
}

func Push(manifest *repo.Manifest, message string, force bool) {
	if Verbose {
		log.Printf("Status")
	}

	pushProjectGroup(&manifest.Private, message, force)
}
