package repo

import "fmt"

func (prj Project) Show() {
	fmt.Printf("RepoUrl: %s\n", prj.RepoUrl)
	fmt.Printf("   Repo: %s\n", prj.Repo)
	fmt.Printf("   Repo spec: %s\n", prj.UseRepo)
	fmt.Printf("   Reference: %s\n", prj.Reference)
	fmt.Printf("   Path: %s\n", prj.Path)
}

func (proj ProjectGroup) Show(grpname string) {
	fmt.Printf("Project Group : %s\n", grpname)
	fmt.Printf("Server : %s\n", proj.Server)
	fmt.Printf("Work Area : %s\n", proj.Workarea)
	fmt.Printf("Default Reference: %s\n", proj.Reference)
	for _, prj := range proj.Projects {
		prj.Show()
	}
}
