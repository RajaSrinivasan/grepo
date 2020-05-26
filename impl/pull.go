package impl

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/RajaSrinivasan/grepo/impl/repo"
)

func buildProject(prj *repo.Project) {
	if Verbose {
		log.Printf("Building in %s", prj.Path)
	}
	if len(prj.Build) < 1 {
		log.Printf("No build command provided. Skipping.")
	}
}

func pullProject(prj *repo.Project) {
	if Verbose {
		log.Printf("Pulling in %s", prj.Path)
	}
	wd, _ := os.Getwd()
	os.Chdir(prj.Path)
	defer os.Chdir(wd)

	cmd := exec.Command("git", "remote", "-v")
	out, err := cmd.Output()
	if err != nil {
		log.Printf("%s", err)
	}
	log.Printf("%s: %s", prj.Path, out)
	if strings.Index(string(out), "not a git repository") > 0 {
		log.Printf("%s: %s", prj.Path, out)
		return
	}
}

func pullProjectGroup(prjg *repo.ProjectGroup, build bool) {

	for _, prj := range prjg.Projects {
		_, err := os.Stat(prj.Path)
		if err != nil {
			log.Printf("%s does not exist. Cannot pull", prj.Path)
			continue
		}
		pullProject(&prj)
		if build {
			buildProject(&prj)
		}
	}
}

func Pull(manifest *repo.Manifest, all bool, build bool) {
	if Verbose {
		log.Printf("Pull. All=%v Build=%v", all, build)
	}
	if all {
		pullProjectGroup(&manifest.Public, build)
	}

	pullProjectGroup(&manifest.Private, build)
}
