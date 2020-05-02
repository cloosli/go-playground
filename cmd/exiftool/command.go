package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gitlab.com/christian.loosli/go-playground/cmd/exiftool/exiftool"
	"gitlab.com/christian.loosli/go-playground/cmd/exiftool/models"
)

func Move() {
	inputDir := config.RootDir
	outputDir := config.OutputDir
	exiftool.Move(inputDir, outputDir)
}

func Check(files *[]models.Exifdata) {
	firstDate := time.Now()
	for _, file := range *files {
		if firstDate.After(*file.CreateDate) {
			firstDate = *file.CreateDate
		}
		if file.DateTimeOriginal.Year() < 2019 {
			fmt.Println("still wrong date", file.String())
		}
	}
	fmt.Println("firstdate", firstDate.String())
	var targetDate time.Time
	if len(config.Args) >= 1 {
		targetDateString := config.Args[0]
		targetDate, _ = time.Parse("20060102T1504", targetDateString)
		fmt.Println("targetDate", targetDate)
		fmt.Println("diff minutes: ", int(targetDate.Sub(firstDate).Round(time.Minute).Minutes()))
	}
}

func Diffs(files *[]models.Exifdata) {
	contains := func(s []int, e int) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}
		return false
	}
	var lookup []int
	for _, file := range *files {
		if !file.GPSDateTime.IsZero() {
			fmt.Println("found gps date")
			diff := file.GPSDateTime.Sub(*file.CreateDate).Round(time.Minute)
			if !contains(lookup, int(diff.Minutes())) {
				lookup = append(lookup, int(diff.Minutes()))
				fmt.Printf("%v diff1 is%v \tdiff1: %d\n", file.FileName, diff.String(), int(diff.Minutes()))
			}
			// file.SetZoneOffset(file.GPSDateTime, config.Offset)
			// diff2 := file.GPSDateTime.UTC().Sub(file.CreateDate.UTC()).Round(time.Minute)
			// diff2 = diff2 + time.Duration(int64(time.Hour)*int64(config.Offset))
			// fmt.Printf("%v diff2 is%v \tdiff2: %d\n", file.FileName, diff2.String(), int(diff2.Minutes()))
		}
	}
}

func Date(files *[]models.Exifdata) {
	for _, file := range *files {
		if !file.GPSDateTime.IsZero() {
			file.SetZoneOffset(file.GPSDateTime, config.Offset)
			if file.DateTimeOriginal.Format("2006:01:02 15:04:05") != file.GPSDateTime.Format("2006:01:02 15:04:05") {
				exiftool.SetAllDates(file.SourceFile, file.GPSDateTime.Format("2006:01:02 15:04:05"))
			}
		} else if !file.CreateDate.IsZero() {
			if len(config.Args) < 1 {
				log.Fatalln("missing argument time shift in minutes")
			}
			minutes, _ := strconv.Atoi(config.Args[0])
			file.SetZoneOffset(file.CreateDate, config.Offset)
			date := *file.CreateDate
			date = date.Add(time.Duration(minutes) * time.Minute)

			if minutes != 0 {
				exiftool.SetAllDates(file.SourceFile, date.Format("2006:01:02 15:04:05"))
			}
		} else {
			log.Fatalln("no date for file", file.String())
		}
		fmt.Println(file.String())
	}
}
