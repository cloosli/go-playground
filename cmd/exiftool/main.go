package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	models "gitlab.com/christian.loosli/go-playground/cmd/exiftool/models"
)

func main() {
	rootDir := "/Users/cloosli/Pictures/GoPro/playground/"

	files, err := getExifdata(rootDir)
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files {
		if i > 1 && i%10 == 0 {
			fmt.Printf(" - - - - - - \n")
		}
		fmt.Printf("%d %v \n", i, f.String())
		if !f.GPSDateTime.IsZero() {
			f.SetZoneOffset(f.GPSDateTime, -7)
			fmt.Printf("  %v \n", f.String())
		}
	}
}

func getExifdata(rootDir string) ([]models.Exifdata, error) {
	checkDir(rootDir)
	s, err := readExiftoolAsJson(rootDir)
	if err != nil {
		return nil, err
	}
	files, err := readJson(s)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func readExiftoolAsJson(filepath string) ([]byte, error) {
	cmd := exec.Command("exiftool", "-r", "-json", filepath)
	var output bytes.Buffer
	cmd.Stdout = &output
	if err := cmd.Start(); err != nil {
		log.Fatalln("Error: ", cmd.String(), err)
	}
	cmd.Wait()
	log.Println("OK ", cmd.String())
	return output.Bytes(), nil
}

func checkDir(filepath string) {
	if _, err := ioutil.ReadDir(filepath); err != nil {
		log.Fatal(err)
	}
}

func readJson(jsonData []byte) ([]models.Exifdata, error) {
	var res []map[string]interface{}
	if err := json.Unmarshal(jsonData, &res); err != nil {
		return nil, err
	}
	list := []models.Exifdata{}
	for _, i := range res {
		var j models.Exifdata
		j.Parse(i)
		list = append(list, j)
	}
	return list, nil
}

func readFile(filepath string) ([]models.Exifdata, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return readJson(data)
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
