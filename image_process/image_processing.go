package imageprocess

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ReadImage(path string) image.Image {
	inputFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println(path)
		panic(err)
	}

	return img
}

func WriteImage(path string, img image.Image) {
	outFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	err = jpeg.Encode(outFile, img, nil)

	if err != nil {
		panic(err)
	}
}

func GrayScale(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Max.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.Gray16Model.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

func Resize(img image.Image) image.Image {
	newHeight := uint(500)
	newWidth := uint(500)

	resizeImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos2)
	return resizeImg
}
