package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var inputFilePath string
	var outputFilePath string
	var isEncode bool
	var isDecode bool

	flag.StringVar(&inputFilePath, "i", "", "Input file to encode or decode")
	flag.StringVar(&outputFilePath, "o", "", "Output file to write encoding or decoding result to")
	flag.BoolVar(&isEncode, "e", true, "Encode QR encoded string (default)")
	flag.BoolVar(&isDecode,"d", false, "Decode QR encoded string")
	flag.Parse()

	if inputFilePath == "" || outputFilePath == "" {
		exit("No input or output file path provided")
	}

	if _, err := os.Stat(inputFilePath); os.IsNotExist(err) {
		exit("Input file does not exist")
	}

	inputBytes, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		exit("File could not be read: " + err.Error())
	}

	var outputBytes []byte
	if isDecode {
		var err error
		outputBytes, err = qrDecode(inputBytes)

		if err != nil {
			exit("Error decoding QR base45 string. Did you forget to escape dollar signs or used single quotes to echo?")
		}
	} else {
		outputBytes = qrEncode(inputBytes)
	}

	ioutil.WriteFile(outputFilePath, outputBytes, 0644)
}

func exit(errorMsg string) {
	fmt.Fprintln(os.Stderr, errorMsg)
	os.Exit(1)
}