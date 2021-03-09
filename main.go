package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	var isEncode bool
	var isDecode bool

	flag.BoolVar(&isEncode, "e", true, "Encode QR encoded string (default)")
	flag.BoolVar(&isDecode,"d", false, "Decode QR encoded string")
	flag.Parse()

	var stdinBuffer bytes.Buffer

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdinBuffer.Write(scanner.Bytes())
	}

	var bytesStdout []byte
	if isDecode {
		var err error
		bytesStdout, err = qrDecode(stdinBuffer.Bytes())

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error decoding QR base45 string. Did you forget to escape dollar signs or used single quotes to echo?")
			os.Exit(1)
		}
	} else {
		bytesStdout = qrEncode(stdinBuffer.Bytes())
	}

	fmt.Println(string(bytesStdout))
}