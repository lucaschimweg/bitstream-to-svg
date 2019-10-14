package lib

import (
	"fmt"
	"io"
)

const svgHeader = `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/2000/svg">
<svg version="1.1"
     xmlns="http://www.w3.org/2000/svg"
     height="80px"
     width="%dpx">

    <defs>
        <style type="text/css">
            text {
                font-size: 15px;
                font-family: monospace;
            }
        </style>
    </defs>
`

const svgFooter = `
</svg>`

const svgTextEntry = `<text x="%d" y="15">%s</text>
`

const svgPrePath = `<path d="`

const svgPostPath = `" stroke="#000" fill="none" stroke-width="2px"/>
`

func writeTextObjects(writer io.Writer, stream BitStream)error {
	xPos := 15
	for _, x := range stream {
		toPr := "0"
		if x {
			toPr = "1"
		}

		_, err := fmt.Fprintf(writer, svgTextEntry, xPos, toPr)
		if err != nil {
			return err
		}
		xPos += 40
	}
	return nil
}

func getY(tk int8)int {
	if tk == 1 {
		return 30
	}
	if tk == 0 {
		return 50
	}
	return 70
}

func writePath(writer io.Writer, converterStream BitStreamConverterStream)error {
	_, err := io.WriteString(writer, svgPrePath)
	if err != nil {
		return err
	}

	if !converterStream.Available() { return nil }
	tk := converterStream.Next()
	fmt.Println(tk)
	posX := 0

	_, err = fmt.Fprintf(writer, "M%d,%d L%d,%d", posX, getY(tk), posX+20, getY(tk))
	if err != nil {
		return err
	}

	posX += 20


	for converterStream.Available() {
		tk = converterStream.Next()
		fmt.Println(tk)
		_, err = fmt.Fprintf(writer, " L%d,%d L%d,%d" , posX, getY(tk), posX+20, getY(tk))
		if err != nil {
			return err
		}

		posX += 20
	}

	_, err = io.WriteString(writer, svgPostPath)
	if err != nil {
		return err
	}

	return nil
}

func WriteSvg(writer io.Writer, stream BitStream, converterStream BitStreamConverterStream)error {
	_, err := fmt.Fprintf(writer, svgHeader, 20 * converterStream.Len())
	if err != nil {
		return err
	}

	err = writeTextObjects(writer, stream)
	if err != nil {
		return err
	}

	err = writePath(writer, converterStream)
	if err != nil {
		return err
	}

	_, err = io.WriteString(writer, svgFooter)
	return err
}
