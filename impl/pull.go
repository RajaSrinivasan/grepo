package impl

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/RajaSrinivasan/grepo/impl/repo"
)

func buildProject(prj *repo.Project) {
	log.Printf("Building in %s", prj.Path)
	if len(prj.Build) < 1 {
		log.Printf("No build command provided. Skipping.")
	}
	flds := strings.Split(prj.Build, " ")
	cmd := exec.Command(flds[0], flds[1:]...)
	if Verbose {
		log.Printf("Executing %s", cmd.String())
	}
	result, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("%s", err)
	}
	if Verbose {
		log.Printf("%s", result)
	}
}

func pullProject(prj *repo.Project) {
	log.Printf("Pulling in %s", prj.Path)

	wd, _ := os.Getwd()
	os.Chdir(prj.Path)
	defer os.Chdir(wd)

	cmd := exec.Command("git", "pull")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("%s", err)
	}
	log.Printf("%s", out)

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
