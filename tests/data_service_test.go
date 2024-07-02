package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/tealeg/xlsx"
)

func TestGetDataForField(t *testing.T) {
    file := loadTestExcelFile()
    data, total, err := GetDataForField("PartnerId", 1, 10)
    
    assert.Nil(t, err)
    assert.Equal(t, 10, len(data))
    assert.True(t, total > 10)
}

func loadTestExcelFile() *xlsx.File {
    file := xlsx.NewFile()
    sheet, _ := file.AddSheet("Sheet1")
    row := sheet.AddRow()
    row.AddCell().Value = "PartnerId"
    row.AddCell().Value = "PartnerName"
    return file
}
