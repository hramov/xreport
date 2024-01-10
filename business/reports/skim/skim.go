package skim

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func Create(data any) ([]byte, error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	return nil, nil
}
