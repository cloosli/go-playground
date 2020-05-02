package exiftool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gitlab.com/christian.loosli/go-playground/cmd/exiftool/models"
)

func ReadExiftoolAsJson(filepath string) ([]byte, error) {
	return do("-r", "-json", filepath)
}

func SetAllDates(filepath string, date string) {
	alldates := fmt.Sprintf("-AllDates='%v'", date)
	datetimeoriginal := fmt.Sprintf("-datetimeoriginal='%v'", date)
	do(alldates, datetimeoriginal, filepath)
}

func UpdateAllDates(filepath string, operand string, seconds int64) {
	allDatesArg := fmt.Sprintf("-AllDates%v='0:0:0 0:0:%d'", operand, seconds)
	do(allDatesArg, filepath)
}

func Rename() {
	// exiftool '-testName<$createdate' -d "IMG_%Y%m%d_%H%M%S%%-c.%%e" -if '$filetype eq "JPEG"' -r .
	// exiftool '-FileName<$createdate' -d "IMG_%Y%m%d_%H%M%S%%-c.%%e" -if '$filetype eq "JPEG"' -r .

	// exiftool '-testName<$createdate' -d "VIDEO_%Y%m%d_%H%M%S%%-c.%%e" -if '$filetype eq "MP4"' -r .
	// exiftool '-FileName<$createdate' -d "VIDEO_%Y%m%d_%H%M%S%%-c.%%e" -if '$filetype eq "MP4"' -r .
}

func Move(inputDir, outputDir string) {
	//exiftool '-Directory<CreateDate' -d newfolder/%Y/%Y-%m-%d -r
	fmt.Printf("move from %v to %v\n", inputDir, outputDir)
	//do("-r", "'-Directory<CreateDate'", "-d", outputDir+"%Y/%Y-%m-%d", inputDir)
}

func GetExifdata(rootDir string) (*[]models.Exifdata, error) {
	checkDir(rootDir)
	s, err := ReadExiftoolAsJson(rootDir)
	if err != nil {
		return nil, err
	}
	files, err := readJson(s)
	if err != nil {
		return nil, err
	}
	return &files, nil
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
