package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	bytesStdin := scanner.Bytes()

	var bytesStdout []byte
	if isDecode {
		var err error
		bytesStdout, err = qrDecode(bytesStdin)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error decoding QR base45 string. Did you forget to escape dollar signs or used single quotes to echo?")
			os.Exit(1)
		}
	} else {
		bytesStdout = qrEncode(bytesStdin)
	}

	fmt.Println(string(bytesStdout))
}