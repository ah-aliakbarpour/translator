package export

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"sort"
	"translator/translate"
)

const SheetName = "Sheet1"

type ExcelExporter struct {
	Data []translate.Result
}

func (exporter *ExcelExporter) Export() error {
	file := excelize.NewFile()

	// write header
	file.MergeCell(SheetName, "B1", "Z1")
	headers := []string{"SOURCE", "TRANSLATIONS"}
	for i, header := range headers {
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	// sort data A-Z
	sort.Slice(exporter.Data, func(i, j int) bool {
		return exporter.Data[i].Source < exporter.Data[j].Source
	})

	// write data
	for i, datum := range exporter.Data {
		file.SetCellValue(SheetName, fmt.Sprintf("A%d", i+2), datum.Source)
		for j, translation := range datum.Translations {
			file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(66+j)), i+2), translation)
			if rune(66+j) == 'Z' {
				break
			}
		}
	}

	// save file
	if err := file.SaveAs("translate.xlsx"); err != nil {
		return fmt.Errorf("can't save the excel file: %w", err)
	}

	return nil
}
