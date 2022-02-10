package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

var (
	of      = flag.String("format", "jpeg", "jpeg | png | gif | bmp | tiff | eps | raw")
	path    = flag.String("path", "", "C:/....")
	outpath = flag.String("out", "", "C:/....")
)

func main() {
	flag.Parse()
	img, err := decode(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to decode provided image:")
	}

	fmt.Println(*outpath, *path)
	f, err := os.Create(fmt.Sprintf("%s.%s", *outpath, *of))
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create file %s: %v", *outpath, err)
	}

	switch *of {
	case "jpeg":
		toJPEG(img, f)
	case "png":
		toPNG(img, f)
	case "gif":
		toGIF(img, f)
	case "bmp":
		toBMP(img, f)
	case "tiff":
		toTIFF(img, f)
	}

}

func decode(in string) (img image.Image, err error) {
	b, err := ioutil.ReadFile(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get file %s: %v\n", in, err)
	}

	r := bytes.NewReader(b)
	img, kind, err := image.Decode(r)
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to decode %s: %v\n", kind, err)
		return
	}
	return
}

func toJPEG(img image.Image, out io.Writer) {
	err := jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	if err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toPNG(img image.Image, out io.Writer) {
	err := png.Encode(out, img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "png: %v\n", err)
		os.Exit(1)
	}
}

func toGIF(img image.Image, out io.Writer) {
	gifOpts := gif.Options{
		NumColors: 256,
	}
	err := gif.Encode(out, img, &gifOpts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gif: %v\n", err)
		os.Exit(1)
	}
}

func toTIFF(img image.Image, out io.Writer) {
	err := tiff.Encode(out, img, &tiff.Options{
		Compression: tiff.Deflate,
		Predictor:   true,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "tiff: %v\n", err)
		os.Exit(1)
	}
}

func toBMP(img image.Image, out io.Writer) {
	err := bmp.Encode(out, img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bmp: %v\n", err)
		os.Exit(1)
	}
}
