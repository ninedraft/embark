package embark

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var packageNameRe = regexp.MustCompile(`^[a-zA-z]+[a-zA-Z]+$`)

type Config struct {
	// flag definitions here
	// https://github.com/octago/sflags#flags-based-on-structures------
	Name           string `flag:"name n" desc:"new package name"`
	Cli            bool   `flag:"cli c" desc:"generate cli boilerplate"`
	Lib            bool   `flag:"lib l" desc:"skip main package generation"`
	PackageManager string `desc:"package manager to use"`
}

func (config Config) Validate() error {
	var errs MultiError
	if strings.TrimSpace(config.Name) == "" {
		errs = append(errs, fmt.Errorf("package name expected to be non empty"))
	}
	if !packageNameRe.MatchString(config.Name) {
		var err = fmt.Errorf("package name must contain only latin chars and '-' (regexp: %q)", packageNameRe)
		errs = append(errs, err)
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}

func (config Config) Dirs() []string {
	var dirs = []string{
		filepath.Join("pkg", config.Name),
	}
	if !config.Lib {
		dirs = append(dirs, filepath.Join("cmd", config.Name))
	}
	return dirs
}

func (config Config) Files() map[string]string {
	var files = make(map[string]string)
	files[filepath.Join("pkg", config.Name, config.Name+".go")] = fmt.Sprintf(`package %s

type Config struct{
	A, B, C string
}

func RenameMe(config Config) {
	fmt.Prinln("Hello, world from " + %q+ "!")
}
`, config.Name, config.Name)

	if !config.Lib {
		var mainCmdFile = fmt.Sprintf(`package main

import "../../pkg/%s"

func main() {
	%s.RenameMe(%s.Config{})
}
`, config.Name, config.Name, config.Name)
		files[filepath.Join("cmd", config.Name, config.Name+".go")] = mainCmdFile
	}
	return files
}

func (config Config) GetPackageManager() (PackageManager, error) {
	switch strings.ToLower(config.PackageManager) {
	case "dep":
		return nil, nil
	case "gomod":
		return nil, nil
	default:
		return nil, fmt.Errorf("invalid package manager %q."+
			"Valid package managers: dep, gomod", config.PackageManager)
	}
}
