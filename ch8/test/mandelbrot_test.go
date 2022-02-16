package test

import (
	"image"
	"image/png"
	"log"
	"os"
	"sync"
	"testing"
)

// 3.2章节代码改造
func TestMandelbrot(t *testing.T) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	wg := sync.WaitGroup{}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			wg.Add(1)
			go func(y float64, px int, py int) {
				defer wg.Done()
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				img.Set(px, py, mandelbrot(z))
			}(y, px, py)
		}
	}
	wg.Wait()

	file, _ := os.OpenFile("/Users/szp/Desktop/Mandelbrot.png", os.O_WRONLY|os.O_CREATE, 0666)

	err := png.Encode(file, img)
	if err != nil {
		log.Fatalln(err)
	}
}
