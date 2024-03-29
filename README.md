# godocc

Like `go doc` but with colors.

## Installation

```
$ go install github.com/inancgumus/godocc@latest
```

## Usage

Accepts all the arguments and flags `go doc` works with. Godocc is just a simple wrapper around the go doc tool.

Example:

```
$ godocc io Writer
```
![godocc sample 1](samples/sample3.png)

```
$ godocc ioutil
```

![godocc sample 2](samples/sample4.png)

## Styling

godocc comes with many colors! Configure the color of the output by setting the following env variable:

```
$ GODOCC_STYLE="dracula"
```

**My favorite styles:** dracula, monokai, fruity, native, paraiso-dark, pygments, rainbow_dash, rrt, solarized-dark, swapoff, vim.

**Other styles:** abap, algol, arduino, autumn, borland, bw, colorful, emacs, friendly, github, igor, lovelace, manni, monokailight, murphy, paraiso-light, pastie, perldoc, solarized-light256, solarized-light, tango, trac, vs, xcode.

_NOTE: Godocc uses the awesome [Chroma](https://github.com/alecthomas/chroma) package underneath._
