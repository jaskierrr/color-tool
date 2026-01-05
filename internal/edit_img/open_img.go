// Package editimg open and decode file by pasted path
package editimg

import (
	"bufio"
	"fmt"
	"image"
	_ "image/png"
	"os"
	"strings"
)

func OpenImg() (image.Image, error) {
	fmt.Print("Введите путь до файла: ")
	reader := bufio.NewReader(os.Stdout)
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	path := strings.TrimSpace(input)
	if path == "1" {path = "/home/jaskier/Pictures/awesome.png"}
	if path == "2" {path = "/home/jaskier/Pictures/first.png"}
	if path == "3" {path = "/home/jaskier/Pictures/mountain.png"}
	fmt.Printf("Image path: %q\n", path)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("Image format: %q\n", format)

	return img, err
}
