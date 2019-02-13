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

	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	var style = "monokai"
	if v := os.Getenv("GODOCC_STYLE"); v != "" {
		style = v
	}

	stdout := colorable.NewColorableStdout()
	err = quick.Highlight(stdout,
		string(output),
		"go", "terminal256", style)
	if err != nil {
		panic(err)
	}

	stdout.Write([]byte("\033[0m"))
}
