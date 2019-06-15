package main

import (
	"log"
	"os"
	"strconv"

	"github.com/disintegration/imaging"
)

func main() {
	argsLen := len(os.Args)
	inputPath := os.Args[argsLen-2]
	src, err := imaging.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	width := src.Bounds().Dx() / 2
	height := src.Bounds().Dy() / 2

	anchors := [...]imaging.Anchor{
		imaging.TopLeft,
		imaging.TopRight,
		imaging.BottomLeft,
		imaging.BottomRight,
	}
	for i, anchor := range anchors {
		img := imaging.CropAnchor(src, width, height, anchor)

		filePrefix := os.Args[argsLen-1]
		err = imaging.Save(img, filePrefix+strconv.Itoa(i+1)+".png")
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
		}
	}
}