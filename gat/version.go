package gat

import (
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

var (
	// appName is the service name for log
	appName = "template"
	// version used for log and info
	appVersion = "unknown"
	goos       = runtime.GOOS
	goarch     = runtime.GOARCH
	gitBranch  = "master"
	gitCommit  = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	buildDate  = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

type version struct {
	// Version is a binary version from git tag.
	Version string `json:"version"`
	// GitCommit is a git commit
	GitCommit string `json:"gitCommit"`
	// GitBranch is a git branch
	GitBranch string `json:"gitBranch"`
	// BuildDate is a build date of the binary.
	BuildDate string `json:"buildDate"`
	// GoOs holds OS name.
	GoOs string `json:"goOs"`
	// GoArch holds architecture name.
	GoArch string `json:"goArch"`
}

// getVersion returns version.
func GetVersion() version {
	return version{
		appVersion,
		gitCommit,
		gitBranch,
		buildDate,
		goos,
		goarch,
	}
}

func GetAppName() string {
	return appName
}

// Print prints version.
func (v version) Print(w io.Writer) {
	fmt.Fprintf(w, "Version: %+v\n", v)
}

func GetNamespace() string {
	return viper.GetString("app.namespace")
}

func GetPhaseEnv() string {
	ns := GetNamespace()
	env := "local"
	strs := strings.Split(ns, "-")
	if len(strs) > 0 {
		env = strs[0]
	}
	return env
}
