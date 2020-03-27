package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sccxx/jpgfixer/"
)

var (
	in  string
	out string
)

func init() {
	flag.StringVar(&in, "i", "", "input")
	flag.StringVar(&out, "o", "output.jpg", "output")
	flag.Parse()
}

func main() {
	if len(in) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	file, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	src, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	dst, err := jpgfixer.Fix(src)
	if err != nil {
		log.Fatal(err)
	}

	outfile, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	outfile.Write(dst)
	outfile.Close()
	fmt.Println("Image fixed ->", out)
}
