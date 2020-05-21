package impl

import (
	"log"
	"os"

	"github.com/RajaSrinivasan/grepo/impl/repo"
)

var Verbose bool

func initProjectGroup(prjg *repo.ProjectGroup, force bool) {
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
			continue
		}
		err = os.MkdirAll(prj.Path, os.ModePerm)
		if err != nil {
			log.Printf("%s", err)
		} else {
			log.Printf("Created %s", prj.Path)
		}
	}
}

func Init(manifest *repo.Manifest, force bool) {
	if Verbose {
		log.Printf("Initializing. Force=%v", force)
	}
	initProjectGroup(&manifest.Public, force)
	initProjectGroup(&manifest.Private, force)
}
