package export

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"sort"
	"translator/dictionary"
)

const Sheet1 = "Sheet1"
const DefaultLastColumn = 'Z'

type Excel struct {
	FileName   string
	LastColumn rune
}

func NewExcel(fileName string, lastColumn rune) Excel {
	if lastColumn < 'B' || lastColumn > 'Z' {
		lastColumn = DefaultLastColumn
	}

	excel := Excel{
		FileName:   fileName,
		LastColumn: lastColumn,
	}

	return excel
}

func (exporter *Excel) Export(data []dictionary.Result) error {
	file := excelize.NewFile()

	// write header
	file.MergeCell(Sheet1, "B1", string(exporter.LastColumn)+"1")
	headers := []string{"SOURCE", "TRANSLATIONS"}
	for i, header := range headers {
		file.SetCellValue(Sheet1, fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}

	// sort data A-Z
	sort.Slice(data, func(i, j int) bool {
		return data[i].Source < data[j].Source
	})

	// write data
	for i, datum := range data {
		file.SetCellValue(Sheet1, fmt.Sprintf("A%d", i+2), datum.Source)
		for j, translation := range datum.Translations {
			file.SetCellValue(Sheet1, fmt.Sprintf("%s%d", string(rune(66+j)), i+2), translation)
			if rune(66+j) == exporter.LastColumn {
				break
			}
		}
	}

	// save file
	if err := file.SaveAs(exporter.FileName + ".xlsx"); err != nil {
		return fmt.Errorf("can't save the excel file: %w", err)
	}

	return nil
}
