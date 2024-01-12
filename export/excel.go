package export

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"translator/translate"
)

const SheetName = "Sheet1"

type ExcelExporter struct {
	Data translate.Result
}

func (exporter *ExcelExporter) Export() error {
	file := excelize.NewFile()

	// write header
	file.MergeCell(SheetName, "B1", "Z1")
	headers := []string{"SOURCE", "TRANSLATIONS"}
	for i, header := range headers {
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	// write body
	i := 2
	for source, translations := range exporter.Data {
		file.SetCellValue(SheetName, fmt.Sprintf("A%d", i), source)
		for j, translation := range translations {
			file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(66+j)), i), translation)
			if i == 26 {
				break
			}
		}
		i++
	}

	// save file
	if err := file.SaveAs("translate.xlsx"); err != nil {
		return fmt.Errorf("can't save the excel file: %w", err)
	}

	return nil
}
