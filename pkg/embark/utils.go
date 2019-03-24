package embark

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

const (
	BashPath = "/usr/bin/bash"
)

type MultiError []error

func MultiErrorMsg(msg string, errs ...error) MultiError {
	var first = MultiError{errors.New(msg)}
	return append(first, errs...)
}

func (errs MultiError) Error() string {
	var buf = &bytes.Buffer{}
	for _, err := range errs {
		fmt.Fprintf(buf, "%v\n", err)
	}
	return buf.String()
}

func GoVersion() (string, error) {
	var buf = &bytes.Buffer{}
	var goVersionErr = RunCmd(buf, "go", "version")
	if goVersionErr != nil {
		return "<undefined>", goVersionErr
	}
	var version = strings.TrimPrefix(buf.String(), "go version")
	return strings.TrimSpace(version), nil
}

func RunCmd(output io.Writer, command string, args ...string) error {
	return (&exec.Cmd{
		Path:   BashPath,
		Stdin:  strings.NewReader(command + " " + strings.Join(args, " ")),
		Stdout: output,
		Stderr: output,
	}).Run()
}
