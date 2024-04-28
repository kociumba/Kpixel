package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	clog "github.com/charmbracelet/log"
	"github.com/ncruces/zenity"
)

func main() {
	sortMethod := flag.String("sort", "", "Sort method: 'column', 'row', 'random'")
	chunkSize = flag.Int("chunk", 10, "chunks to devide the image in to when using random sort")
	flag.Parse()
	if !flag.Parsed() {
		clog.Fatal("Please specify a sorting method")
		flag.Usage()
	}

	var imgPath string
	clog.Info(os.Args)
	imgPath = os.Args[len(os.Args)-1]
	_, err := os.Stat(imgPath)
	if err == nil {
		// fmt.Println("The last argument is a valid filepath:", imgPath)
		// Proceed with using the file
		imgPath = filepath.Clean(imgPath)
	} else {
		clog.Info("Pick the image you want to sort")
		imgPath, _ = zenity.SelectFile(
			zenity.Filename(os.ExpandEnv("$HOME")),
			zenity.FileFilters{{Name: "Images", Patterns: []string{"*.png", "*.jpg", "*.jpeg"}}},
		)
	}
	// if err != nil {
	// 	clog.Fatal(err)
	// }

	img, format, err := openAndDecodeImage(filepath.Clean(imgPath))
	if err != nil {
		clog.Fatal(err)
	}

	strategy, err := parseSortingMethod(*sortMethod)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	newImg := strategy.SortImage(img)
	outputPath := filepath.Dir(imgPath) + "/" + filepath.Base(imgPath) + ".sorted." + format
	if err := saveImage(outputPath, newImg, format); err != nil {
		clog.Fatal(err)
	}

	clog.Info("Image sorted and saved successfully.")
}
