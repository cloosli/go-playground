package main

import "testing"
import "os"
import "fmt"

func Test_check_error_with_nil(t *testing.T) {
	checkError("Message", nil)
}

func Test_check_error(t *testing.T) {
	err := fmt.Errorf("Ignore")
	checkError("Message", err)
	t.Failed()
}

func Test_write_data_to_file(t *testing.T) {
	file, err := os.NewFile(544, "result.csv")

	main()

	t.Error("failed")
}
