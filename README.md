# bitstream-to-svg
With this tool, you can create SVGs from bitstreams, to show the possible encodings NZR, NZRI, AMI and Manchester.

The tool is implemented in Go and can be used with a simple command line interface.

It is built with GoModules. If you have Go installed, you can just clone this repo and hit `go build .`. This will build the project and move the executable into the current working dir. 

## Command Line Interface
```
Usage: ./bitstream-to-svg [flags] bitstream
  -e string
       Encoding to use (nrz, nrzi, manchester, ami)
  -f string
       SVG Output File (default "out.svg")
```
So one example call would be:

```./bitstream-to-svg -e ami -f out.svg 0010101```

This will write the output SVG to `out.svg`.
