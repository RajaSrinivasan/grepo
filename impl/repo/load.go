package repo

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var project Manifest

func LoadConfig(cfgfile string) {

	if Verbose {
		log.Printf("Loading the config file %s", cfgfile)
	}

	yamlin, _ := os.Open(cfgfile)
	defer yamlin.Close()

	yamlobjin := yaml.NewDecoder(yamlin)
	yamlobjin.Decode(&project)

	project.Public.Show("Public")
	project.Private.Show("Private")

}
