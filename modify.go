package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//	inp := flag.String("input", "midi1", "midi input device in the /dev/ folder")
	out := flag.String("output", "midi1", "midi output device in the /dev/ folder")
	flag.Parse()

	//	inDev := fmt.Sprintf("/dev/%s", *inp)
	//	fmt.Printf("Input device: %v\n", inDev)

	outDev := fmt.Sprintf("/dev/%s", *out)
	fmt.Printf("Output device: %v\n", outDev)

	//	fIn, err := os.Open(inDev)
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer fIn.Close()

	fOut, err := os.OpenFile(outDev, os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer fOut.Close()

	for i := 0; i < 1000; i++ {
		//	nByte := readNextBytes(fIn, 3)
		nByte := readNextBytes(fOut, 3)
		if nByte[0] != 0xf8 {
			nByte[0] += 13
			fmt.Print(hex.EncodeToString(nByte))

			fmt.Println(nByte)

			writeNextBytes(fOut, nByte)
		}
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func writeNextBytes(file *os.File, bytes []byte) {

	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}

}
