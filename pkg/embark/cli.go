package embark

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	gpflag "github.com/octago/sflags/gen/gpflag"
	cobra "github.com/spf13/cobra"
)

func Main() *cobra.Command {
	var config = Config{
		Lib: true,
	}
	var cmd = &cobra.Command{
		Use: "embark",
		Run: func(cmd *cobra.Command, args []string) {
			if err := config.Validate(); err != nil {
				ErrPrintln("error:")
				ErrPrintln(err)
				os.Exit(100)
			}
			fmt.Printf("Creating dirs:\n- %s\n", strings.Join(config.Dirs(), "\n- "))
			if err := CreateDirs("", config.Dirs()...); err != nil {
				ErrPrintln(err)
				os.Exit(100)
			}
			var files = config.Files()
			fmt.Println("Creating files:")
			for filename := range files {
				fmt.Printf(" - %s\n", filename)
			}
			if err := CreateFiles("", files); err != nil {
				ErrPrintln(err)
				os.Exit(10)
			}
			if err := RunCmd(os.Stdout, "git", "init"); err != nil {
				ErrPrintln(err)
				os.Exit(100)
			}
			if err := RunCmd(os.Stdout, "dep", "init", "-v"); err != nil {
				ErrPrintln(err)
				os.Exit(100)
			}
		},
	}
	if err := gpflag.ParseTo(&config, cmd.PersistentFlags()); err != nil {
		panic(err)
	}
	return cmd
}

func CreateDirs(root string, dirs ...string) error {
	var errs MultiError
	for _, dirName := range dirs {
		var fullDirPath = filepath.Join(root, dirName)
		if err := os.MkdirAll(fullDirPath, os.ModeDir|os.ModePerm); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}

func CreateFiles(root string, files map[string]string) error {
	var errs MultiError
	for filename, contents := range files {
		var data = []byte(contents)
		var tergetPath = filepath.Join(root, filename)
		if err := ioutil.WriteFile(tergetPath, data, os.ModePerm); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}
