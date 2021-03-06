package main

import (
	"os"
	"os/exec"
)

func collect(verbose bool, corePath string) error {
	cmd := exec.Command("/usr/libexec/platform-python",
		"-m",
		"insights.collect",
		"--compress")
	if verbose {
		cmd.Args = append(cmd.Args, "--verbose")
	}
	cmd.Env = []string{
		"PATH=" + os.Getenv("PATH"),
		"LANG=" + os.Getenv("LANG"),
		"PYTHONPATH=" + corePath,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return nil
}
