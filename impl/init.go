package impl

import (
	"log"
	"os"
	"os/exec"

	"github.com/RajaSrinivasan/grepo/impl/repo"
)

var Verbose bool

func initProjectGroup(prjg *repo.ProjectGroup, force bool, detached bool) {
	if force {
		err := os.RemoveAll(prjg.Workarea)
		if err != nil {
			log.Printf("%s", err)
			return
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

		wd, _ := os.Getwd()
		os.Chdir(prj.Path)
		defer os.Chdir(wd)

		cmd := exec.Command("git", "init")
		result, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("%s", err)
		} else {
			if Verbose {
				log.Printf(string(result))
			}
		}

		switch detached {
		case true:
			cmd := exec.Command("git", "clone", prj.UseRepo, "-b", prj.Reference)
			result, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("%s\n%s", err, result)
			} else {
				if Verbose {
					log.Printf("%s", result)
				}
			}
		case false:
			cmd := exec.Command("git", "clone", prj.UseRepo)
			result, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("%s\n%s", err, result)
			} else {
				if Verbose {
					log.Printf("%s", result)
				}
				cmd2 := exec.Command("git", "checkout", prj.Reference)
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
	}
}

func Init(manifest *repo.Manifest, force bool) {
	if Verbose {
		log.Printf("Initializing. Force=%v", force)
	}
	initProjectGroup(&manifest.Public, force, true)
	initProjectGroup(&manifest.Private, force, false)
}
