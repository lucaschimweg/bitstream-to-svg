package main

import (
	"flag"
	"fmt"
	"github.com/lucaschimweg/bitstream-to-svg/lib"
	"os"
)

func printHelp() {
	fmt.Printf("Usage: %s [flags] bitstream\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	var encoding = flag.String("e", "", "Encoding to use (nrz, nrzi, manchester, ami)")
	var out = flag.String("f", "out.svg", "SVG Output File")

	flag.Parse()

	if *encoding == "" {
		fmt.Println("Error: You have to specify the encoding!")
		printHelp()
		os.Exit(-1)
	}

	if flag.Arg(0) == "" {
		fmt.Println("Error: You have to specify the bitstream!")
		printHelp()
		os.Exit(-1)
	}

	x, err := lib.StringToBitStream(flag.Arg(0))

	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(-1)
	}

	y, err := lib.CreateConverterStream(*encoding, x)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		printHelp()
		os.Exit(-1)
	}

	filename := *out

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = file.Close()
	}()

	err = lib.WriteSvg(file, x, y)

	if err != nil {
		fmt.Println(err)
	}

}
