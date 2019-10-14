package main

import (
	"fmt"
	"github.com/lucaschimweg/bitstream-to-svg/lib"
	"os"
)

func main() {
	x, err := lib.StringToBitStream("10000101111")

	if err != nil {
		panic(err.Error())
	}

	filename := "output.svg"

	file, err := os.Create(filename)


	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		_ = file.Close()
	}()

	y := lib.CreateAmiConverterStream(x)
	err = lib.WriteSvg(file, x, y)

	if err != nil {
		fmt.Println(err)
	}

}
