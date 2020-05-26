package repo

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var project Manifest

func (prj *Project) fixupProject(grp *ProjectGroup) error {
	log.Printf("Fixing up %s (%s)", prj.Repo, prj.RepoUrl)
	var userepo string
	if len(prj.RepoUrl) > 1 {
		userepo = prj.RepoUrl
	} else {
		userepo = grp.Server + prj.Repo
	}
	prj.UseRepo = userepo

	if len(prj.Path) < 1 {
		log.Printf("Always provide a work area for repo %s", userepo)
		return errors.New("No path provided for " + userepo)
	}

	prj.Path = filepath.Join(grp.Workarea, prj.Path)

	if len(prj.Reference) < 1 {
		if len(grp.Reference) > 0 {
			prj.Reference = grp.Reference
		}
	}

	return nil
}

func (prjg *ProjectGroup) fixupGroup(nm string) error {
	log.Printf("Fixing up %s", nm)
	if len(prjg.Workarea) < 2 {
		log.Printf("ProjGroup: %s requires a work area", nm)
		return errors.New("No workspace provided for " + nm)
	}
	ws, err := filepath.Abs(prjg.Workarea)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	prjg.Workarea = ws

	for idx, _ := range prjg.Projects {
		log.Printf("Project No : %d", idx)
		err := (&prjg.Projects[idx]).fixupProject(prjg)
		if err != nil {
			log.Printf("%s", err)
		}
	}
	return nil
}

func LoadConfig(cfgfile string) *Manifest {

	if Verbose {
		log.Printf("Loading the config file %s", cfgfile)
	}

	yamlin, _ := os.Open(cfgfile)
	defer yamlin.Close()

	yamlobjin := yaml.NewDecoder(yamlin)
	yamlobjin.Decode(&project)
	(&project.Public).fixupGroup("Public")
	(&project.Private).fixupGroup("Private")

	if Verbose {
		project.Public.Show("Public")
		project.Private.Show("Private")
	}
	return &project
}
