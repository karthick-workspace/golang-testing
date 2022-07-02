package example

import (
	"fmt"
	"io"
)

var file string = "this is not used"

func Example_crop() {
	var r io.Reader

	img, err := Decode(r)
	if err != nil {
		panic(err)
	}

	err = Crop(img, 0, 0, 20, 20)
	if err != nil {
		panic(err)
	}

	var w io.Writer
	if err != nil {
		panic(err)
	}

	err = Encode(img, w)
	if err != nil {
		panic(err)
	}

	fmt.Println("See out.jpg for the cropped image")

	// Output:
	// See out.jpg for the cropped image
}
