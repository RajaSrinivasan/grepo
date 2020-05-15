package version

import "fmt"

func Report() {
	fmt.Printf("Version : %d.%d-%d\n", versionMajor, versionMinor, versionBuild)
	fmt.Printf("Built : %s\n", buildTime)
	fmt.Printf("Repo URL : %s\n", repoURL)
	fmt.Printf("Branch : %s\n", branchName)
	fmt.Printf("Commit Id : Short : %s Long %s\n", shortCommitId, longCommitId)
	fmt.Printf("Assigned Tags : %s\n", assignedTags)

}
