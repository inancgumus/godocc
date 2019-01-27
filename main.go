package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alecthomas/chroma/quick"
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

	err = quick.Highlight(os.Stdout,
		string(output),
		"go", "terminal256", style)
	if err != nil {
		panic(err)
	}

	fmt.Print("\033[0m")
}
