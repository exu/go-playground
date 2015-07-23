package main

import (
	"bytes"
	"code.google.com/p/rsc/qr"
	"crypto/rand"
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"
)

func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	if randType == "number" {
		dictionary = "0123456789"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func main() {

	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// generate a random string - preferably 6 or 8 characters
	randomStr := randStr(6, "alphanum")
	fmt.Println(randomStr)

	// generate the link or any data you want to
	// encode into QR codes
	// in this example, we use an example of making purchase by QR code.
	// Replace the stuff2buy with yours.

	stuff2buy := "stuffpurchaseby?userid=" + randomStr + "&issuer=SomeBigSuperMarket"

	// Encode stuff2buy to QR codes
	// qr.H = 65% redundant level
	// see https://godoc.org/code.google.com/p/rsc/qr#Level

	code, err := qr.Encode(stuff2buy, qr.H)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imgByte := code.PNG()

	// convert byte to image for saving to file
	img, _, _ := image.Decode(bytes.NewReader(imgByte))

	//save the imgByte to file
	out, err := os.Create("./QRImg.png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, img)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// everything ok
	fmt.Println("QR code generated and saved to QRimg1.png")

}
