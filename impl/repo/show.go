package repo

import "fmt"

func (prj Project) Show() {
	fmt.Printf("RepoUrl: %s\n", prj.RepoUrl)
	fmt.Printf("   Repo: %s\n", prj.Repo)
	fmt.Printf("   Reference: %s\n", prj.Reference)
	fmt.Printf("   Path: %s\n", prj.Path)
}

func (proj ProjectGroup) Show(grpname string) {
	fmt.Printf("Project Group %s\n", grpname)
	fmt.Printf("Server [default] %s\n", proj.Server)
	fmt.Printf("Protocol [default] : %s", proj.Protocol)
	fmt.Printf("Work Area : %s\n", proj.Workarea)
	for _, prj := range proj.Projects {
		prj.Show()
	}
}