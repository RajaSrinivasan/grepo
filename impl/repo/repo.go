package repo

var Verbose bool

type Project struct {
	RepoUrl   string // This is the complete url including protocol. If not specified group spec
	Repo      string // One of repourl or repo can be specified. repo => rest from group spec
	UseRepo   string // final derivation which is actually used
	Reference string // Branch, tag or commit id. if not specified the "default" as setup in the repo to read
	Path      string // The folder where the checkout is performed under the Work area
	Build     string // Build Command. Execute this shell command upon pull
}

type Projects []Project

type ProjectGroup struct {
	Workarea  string // Local work area where checked out
	Server    string // Server name e.g. github.com | gitlab.com. Default for all Projects
	Reference string // Default branch/tag
	Projects  Projects
}

type Manifest struct {
	Public  ProjectGroup // public. these are generally not touched by developer
	Private ProjectGroup // repos touched by developer
}

const NoneBranch = "<none>"
const NoneTag = "<none>"
