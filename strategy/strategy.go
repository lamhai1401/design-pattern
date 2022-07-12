package strategy

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"log"
	"os"
)

type PrintStrategy interface {
	SetLog(io.Writer)
	SetWriter(io.Writer)
	Print() error
}

type ConsoleSquare struct {
	PrintOutput
}

type ImageSquare struct {
	DestinationFilePath string
	PrintOutput
}

func (c *ConsoleSquare) Print() error {
	println("Square")
	return nil
}

func (t *ImageSquare) Print() error {
	width := 800
	height := 600
	origin := image.Point{0, 0}
	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	quality := &jpeg.Options{Quality: 75}
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	w, err := os.Create(t.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("error opening image")
	}
	defer w.Close()
	if err = jpeg.Encode(w, bgImage, quality); err != nil {
		return fmt.Errorf("error writing image to disk")
	}
	return nil
}

func main() {

	var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

	activeStrategy, err := NewPrinter(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case IMAGE_STRATEGY:
		w, err := os.Create("/tmp/image.jpg")
		if err != nil {
			log.Fatal("Error opening image")
		}
		defer w.Close()
		activeStrategy.SetWriter(w)
	}

	err = activeStrategy.Print()
	if err != nil {
		log.Fatal(err)
	}
}

type TextSquare struct {
	PrintOutput
}

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (t *TextSquare) Print() error {
	r := bytes.NewReader([]byte("Circle"))
	io.Copy(t.Writer, r)
	return nil
}

func (d *PrintOutput) SetLog(w io.Writer) {
	d.LogWriter = w
}
func (d *PrintOutput) SetWriter(w io.Writer) {
	d.Writer = w
}

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

func NewPrinter(s string) (PrintStrategy, error) {
	switch s {
	case TEXT_STRATEGY:
		return &TextSquare{
			PrintOutput: PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			PrintOutput: PrintOutput{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("strategy '%v' not found", s)
	}
}
