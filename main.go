package main

import (
	"os"
	"os/exec"

	"github.com/alecthomas/chroma/quick"
	colorable "github.com/mattn/go-colorable"
)

func main() {
	// capture the input, output and args of `go doc`.
	args := []string{"doc"}
	args = append(args, os.Args[1:]...)
	cmd := exec.Command("go", args...)
	cmd.Env = os.Environ()

	// ignore the error for now.
	// output contains the error as well.
	// so the error will be printed eventually.
	output, _ := cmd.CombinedOutput()

	// remove the newline first.
	// reprint it later on: after turning off the coloring.
	if output[len(output)-1] == '\n' {
		output = output[:len(output)-1]
	}

	// default style is dracula
	var style = "dracula"
	if v := os.Getenv("GODOCC_STYLE"); v != "" {
		style = v
	}

	// get a cross-platform bash escape sequences emulator
	stdout := colorable.NewColorableStdout()
	
	// colorize and print the documentation
	if err := quick.Highlight(stdout, string(output), "go", "terminal256", style); err != nil {
		panic(err)
	}

	// turn off the coloring and reprint a newline
	stdout.Write([]byte("\033[0m\n"))
}
