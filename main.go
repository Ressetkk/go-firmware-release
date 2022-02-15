package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"github.com/Ressetkk/go-firmware-release/fw"
	"io"

	"os"
	"path/filepath"
	"time"
)

type config struct {
	erase              bool
	outName, extension string
}

const (
	WelcomeMsg = "firmware-release started at %s\nInput file: %s\nOutput file: %s\n"
)

func main() {
	c := config{}
	flag.BoolVar(&c.erase, "erase-eeprom", true, "Create file that erases EEPROM.")
	flag.StringVar(&c.outName, "custom-name", "", "Set custom output file name.")
	flag.StringVar(&c.extension, "extension", "bin", "Define to which the file should be saved.")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Error: Provide one input file.")
		flag.Usage()
		os.Exit(1)
	}
	input := filepath.FromSlash(args[0])
	i, err := os.Open(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	p := fw.Packer{
		Reader:  i,
		MaxSize: 0x78000,
		Erase:   c.erase,
	}

	outfile := filename(c)
	o, err := os.Create(outfile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	m := fw.Checksum{
		Writer: o,
		Hash:   md5.New(),
	}

	fmt.Printf(WelcomeMsg, time.Now().UTC(), i.Name(), o.Name())
	if _, err := io.Copy(&m, &p); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Checksum: %x\n", m.Hash.Sum(nil))
	fmt.Println("Success!")
}

func filename(c config) string {
	t := time.Now()
	var out string
	if len(c.outName) > 0 {
		return c.outName
	}
	out = fmt.Sprintf("main_board_%v", t.Format("20060102"))
	if c.erase {
		out += "_erase_eeprom"
	}
	out += "." + c.extension
	return out
}
