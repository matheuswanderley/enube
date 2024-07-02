package services

import (
    "github.com/tealeg/xlsx"
    "strings"
    "fmt"
)

func normalizeField(input string) string {
    return strings.TrimSpace(strings.ToLower(input))
}

func GetDataForField(fieldName string, page int, limit int) ([]string, int, error) {
    file, err := xlsx.OpenFile("dados.xlsx")
    if err != nil {
        return nil, 0, err
    }

    sheet := file.Sheets[0]
    normalizedFieldName := normalizeField(fieldName)

    columnIndex := -1
    for idx, cell := range sheet.Rows[0].Cells {
        if normalizeField(cell.String()) == normalizedFieldName {
            columnIndex = idx
            break
        }
    }

    if columnIndex == -1 {
        return nil, 0, fmt.Errorf("field '%s' not found", fieldName)
    }

    var results []string
    start := (page - 1) * limit
    end := start + limit

    for i, row := range sheet.Rows[1:] {
        if i >= start && i < end {
            if len(row.Cells) > columnIndex {
                results = append(results, row.Cells[columnIndex].String())
            }
        }
    }

    total := len(sheet.Rows) - 1
    return results, total, nil
}
