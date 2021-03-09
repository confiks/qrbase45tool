package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytesStdin, _ := reader.ReadBytes('\n')

	qrEncodedStdin := qrEncode(bytesStdin)
	fmt.Println(string(qrEncodedStdin))
}