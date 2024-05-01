package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	clog "github.com/charmbracelet/log"
	"github.com/ncruces/zenity"
)

var chunkSize *int

func main() {

	// profiling for debug
	// mux := http.NewServeMux()
	// statsviz.Register(mux)

	// go func() {
	// 	clog.Info(http.ListenAndServe("localhost:8080", mux))
	// }()

	// if runtime.GOOS == "windows" {
	// 	checkRegistry() // checks for the registry keys on windows to add a contex menu item to images
	// }

	sortMethod := flag.String("sort", "", "Sort method: 'column', 'row', 'random'")
	chunkSize = flag.Int("chunk", 10, "chunks to devide the image in to when using random sort")
	sortValue := flag.String("method", "hue", "Sort method: 'hue', 'luminosity', 'saturation', 'red', 'green', 'blue'")
	flag.Parse()
	if !flag.Parsed() {
		clog.Fatal("Please specify a sorting method")
		flag.Usage()
	}

	// clog.Info(*chunkSize)

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

	strategy, err := setupSorting(*sortMethod, *sortValue)
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
