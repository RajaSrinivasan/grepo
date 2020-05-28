package impl

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/RajaSrinivasan/grepo/impl/repo"
)

var Verbose bool

func initProjectGroup(prjg *repo.ProjectGroup, force bool, build bool, detached bool, branch string) {
	if force {
		err := os.RemoveAll(prjg.Workarea)
		if err != nil {
			log.Printf("%s", err)
			return
		}
		if Verbose {
			log.Printf("Removed directory tree %s", prjg.Workarea)
		}
	}

	_, err := os.Stat(prjg.Workarea)
	if err == nil {
		log.Printf("%s already exists", prjg.Workarea)
	} else {
		err := os.MkdirAll(prjg.Workarea, os.ModePerm)
		if err != nil {
			log.Printf("%s", err)
			return
		}
	}
	wd, _ := os.Getwd()
	defer os.Chdir(wd)

	for _, prj := range prjg.Projects {
		_, err := os.Stat(prj.Path)
		if err == nil {
			log.Printf("%s already exists", prj.Path)
		} else {
			err = os.MkdirAll(prj.Path, os.ModePerm)
			if err != nil {
				log.Printf("%s", err)
			} else {
				log.Printf("Created %s", prj.Path)
			}
		}
		os.Chdir(prj.Path)

		switch detached {
		case true:
			cmd := exec.Command("git", "clone", prj.UseRepo, "-b", prj.Reference, ".")
			if Verbose {
				log.Printf("Executing %s", cmd.String())
			}
			result, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("%s\n%s", err, result)
				continue
			} else {
				if Verbose {
					log.Printf("%s", result)
				}
			}
		case false:
			cmd := exec.Command("git", "clone", prj.UseRepo, ".")
			if Verbose {
				log.Printf("Executing %s", cmd.String())
			}
			result, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("%s\n%s", err, result)
				continue
			} else {
				if Verbose {
					log.Printf("%s", result)
				}
				cmd2 := exec.Command("git", "checkout", prj.Reference)
				if Verbose {
					log.Printf("Executing %s", cmd2.String())
				}
				result, err := cmd2.CombinedOutput()
				if err != nil {
					log.Printf("%s\n%s", err, result)
				} else {
					if Verbose {
						log.Printf("%s", result)
					}
				}
			}
		}
		if strings.Compare(branch, repo.NoneBranch) != 0 {
			cmd := exec.Command("git", "checkout", "-b", branch)
			if Verbose {
				log.Printf("Executing %s", cmd.String())
			}
			result, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("%s", err)
			}
			log.Printf("%s", result)
		}
		if build {
			if len(prj.Build) < 1 {
				log.Printf("No build instruction provided. Skipping")
				continue
			}
			flds := strings.Split(prj.Build, " ")
			cmd := exec.Command(flds[0], flds[1:]...)
			if Verbose {
				log.Printf("Executing %s", cmd.String())
			}
			result, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("%s\n%s", err, result)
				continue
			}
			if Verbose {
				log.Printf("%s", result)
			}
		}

	}
}

func Init(manifest *repo.Manifest, force bool, build bool, branch string) {
	if Verbose {
		log.Printf("Initializing. Force=%v", force)
	}
	initProjectGroup(&manifest.Public, force, build, true, repo.NoneBranch)
	initProjectGroup(&manifest.Private, force, build, false, branch)
}
