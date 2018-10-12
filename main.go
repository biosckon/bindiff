package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("expecting at least 2 arguments")
		return
	}

	nfi := len(os.Args) - 1

	bufs := make([][]byte, nfi)

	// read in the files
	for i, fn := range os.Args[1:] {
		fi, err := os.Open(fn)
		if err != nil {
			fmt.Printf("error opening file %s: %v", fn, err)
			return
		}
		defer fi.Close()

		bufs[i], err = ioutil.ReadAll(fi)
		if err != nil {
			fmt.Printf("error reading from file %s: %v", fn, err)
		}
	}

	// get a minimum len of all the buffers
	min := len(bufs[0])
	for _, b := range bufs[1:] {
		l := len(b)
		if l < min {
			min = l
		}
	}

	// compare values one by one
	for i := 0; i < min; i++ { // loop over offset

		v := bufs[0][i]            // a byte from first file
		for j := 1; j < nfi; j++ { // loop over the rest of files
			if v != bufs[j][i] {
				fmt.Printf("first difference at offset %d = 0x%x\n", i, i)
				return
			}
		}
	}

	fmt.Println("no differences found")
}
