package main

import (
	"os"
	"os/exec"

	"github.com/alecthomas/chroma/quick"
	colorable "github.com/mattn/go-colorable"
)

func main() {
	args := []string{"doc"}
	args = append(args, os.Args[1:]...)
	cmd := exec.Command("go", args...)
	cmd.Env = os.Environ()

	// ignore the error for now.
	// output contains the error as well.
	// so the error will be printed.
	output, _ := cmd.CombinedOutput()

	// remove the newline first
	// then reprint later on: after turning off the coloring
	if output[len(output)-1] == '\n' {
		output = output[:len(output)-1]
	}

	var style = "monokai"
	if v := os.Getenv("GODOCC_STYLE"); v != "" {
		style = v
	}

	stdout := colorable.NewColorableStdout()
	err := quick.Highlight(stdout, string(output),
		"go", "terminal256", style)
	if err != nil {
		panic(err)
	}

	// turn off the coloring and reprint a newline
	stdout.Write([]byte("\033[0m\n"))
}
