package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type Pixel struct {
	R, G, B, A uint8
}

func createZeroIntSlice(pixels [][]int) (result [][]int) {
	result = make([][]int, len(pixels))
	for i := 0; i < len(pixels); i++ {
		result[i] = make([]int, len(pixels[i]))
	}
	return result
}

func createZeroPixelSlice(pixels [][]Pixel) (result [][]Pixel) {
	result = make([][]Pixel, len(pixels))
	for i := 0; i < len(pixels); i++ {
		result[i] = make([]Pixel, len(pixels[i]))
	}
	return result
}

// Task 2.a
func RotateImage(pixels [][]int) [][]int {
	if len(pixels) < 1 {
		return nil
	}

	col := len(pixels[0]) - 1
	result := createZeroIntSlice(pixels)

	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[0]); j++ {
			result[j][col] = pixels[i][j]
		}
		col--
	}
	return result
}

// Task 2.b
// To test this function ,you can execute the RunRotateActualImage() on main
func RotateActualImage(pixels [][]Pixel) [][]Pixel {
	col := len(pixels[0]) - 1
	result := createZeroPixelSlice(pixels)

	for i := 0; i < len(pixels); i++ {
		for j := 0; j < len(pixels[0]); j++ {
			result[j][col] = pixels[i][j]
		}
		col--
	}
	return result
	// --------------------
}

// Please don't modify code below unless specified by your trainers

const originalImageFile = "./img/img.png"
const rotatedImageFile = "./result/rotated-img.png"

// Below are codes to make this exercise more realistic
// The process can be more straightforward
// However, since we want to balance the complexity
// We mainly divide the steps into 3:
// 1. Extract the pixels
// 2. Rotate the pixels (which you need to implement)
// 3. Create a new image from the rotated pixels
//
// If you are interested, you can simply challenge yourself to
// rotate the image immediately instead of doing these 3 steps,
// but it is not in the scope of this assignment
func RunRotateActualImage() {
	// Open the image file
	file, err := os.Open(originalImageFile)
	if err != nil {
		fmt.Println("Failed to open image:", err)
		return
	}
	defer file.Close()

	// Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Println("Failed to decode image:", err)
		return
	}

	// Manipulate the image (invert colors)
	rotImg := rotate90DegreesRight(img)

	// Save the manipulated image
	outputFile, err := os.Create(rotatedImageFile)
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		return
	}
	defer outputFile.Close()

	switch strings.ToLower(format) {
	case "png":
		png.Encode(outputFile, rotImg)
	default:
		fmt.Println("Unsupported format:", format)
	}
}

// Extracting the pixels
func extractPixels(src image.Image) [][]Pixel {
	bounds := src.Bounds()
	pixels := make([][]Pixel, bounds.Dy())
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		row := make([]Pixel, bounds.Dx())
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := src.At(x, y).RGBA()
			row[x-bounds.Min.X] = Pixel{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}
		}
		pixels[y-bounds.Min.Y] = row
	}
	return pixels
}

// From the rotated pixels, we create a new image
func createImageFromPixels(pixels [][]Pixel) *image.RGBA {
	height, width := len(pixels), len(pixels[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := pixels[y][x]
			img.Set(x, y, color.RGBA{p.R, p.G, p.B, p.A})
		}
	}
	return img
}

func rotate90DegreesRight(src image.Image) *image.RGBA {
	pixels := extractPixels(src)
	rotatedPixels := RotateActualImage(pixels)
	return createImageFromPixels(rotatedPixels)
}
