package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"gitlab.com/christian.loosli/go-playground/cmd/exiftool/exiftool"
)

type CmdConfig struct {
	RootDir   string
	OutputDir string
	Cmd       string
	Offset    int
	Args      []string
}

var config CmdConfig

func main() {
	config = *checkFlags()
	files, err := exiftool.GetExifdata(config.RootDir)
	checkError(err)
	switch config.Cmd {
	case "check":
		Check(files)
	case "diff":
		Diffs(files)
	case "date":
		Date(files)
	case "move":
		Move()
	}
}

func checkFlags() *CmdConfig {
	config := CmdConfig{}
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()

	flag.StringVar(&config.RootDir, "i", "./input", "read directory")
	flag.StringVar(&config.OutputDir, "d", "./output", "path to new pictures directory")
	flag.IntVar(&config.Offset, "offset", 0, "offset")

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fatalErr = errors.New("invalid usage; must specify command")
		return nil
	}
	config.Cmd = args[0]
	config.Args = args[1:]
	fmt.Printf("%+v\n", config)
	return &config
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func walk() {
// files, err := ioutil.ReadDir(rootDir)
// if err != nil {
// 	log.Fatal(err)
// }
// // for _, file := range files {
// // 	log.Println(file.Name())
// // }
// err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// })
// if err != nil {
// 	log.Fatal(err)
// }
// files, err := readFile(path.Join(rootDir, "/guerreronegro/exiftool.json"))
// }
