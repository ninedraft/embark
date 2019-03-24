package embark

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
)

var _ PackageManager = &Dep{}

type Dep struct {
	Version string
	Root    string
}

func InitDep(root string) (*Dep, error) {
	var output = &bytes.Buffer{}
	var depCmd = &exec.Cmd{
		Path:   BashPath,
		Stdin:  strings.NewReader("dep version"),
		Stdout: output,
	}
	if err := depCmd.Run(); err != nil {
		return nil, err
	}
	var dep = &Dep{}
	for _, line := range strings.Split(output.String(), "\n") {
		var tokens = strings.SplitN(line, ":", 2)
		if len(tokens) == 2 {
			var key = strings.TrimSpace(tokens[0])
			if key == "version" {
				var version = strings.TrimSpace(tokens[1])
				dep.Version = version
			}
		}
	}
	return dep, nil
}

func (dep *Dep) String() string {
	return "dep " + dep.Version
}

func (dep *Dep) InitPackage(log io.Writer) error {
	var depCmd = &exec.Cmd{
		Path:   "dep",
		Args:   []string{"init", "-v"},
		Stdout: log,
		Stderr: log,
	}
	return depCmd.Run()
}
