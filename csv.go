package helpers


import (
	"encoding/csv"
	"fmt"
	"io"
)


type record struct {
	headerIndexMap *map[string]int
	values []string
}



func (r *record) Lookup(header string) string {
	headerIndex, found := (*r.headerIndexMap)[header]
	if !found {
		panic("header title not found")
	}
	return r.values[headerIndex]
}

func CSVReaderToRecords(csvReader *csv.Reader) []record {
	
	headerSlice, err := csvReader.Read()
	
	if err != nil {
		panic(fmt.Sprintf("Error reading header with %v", err))
	}

	headerIndexMap := GetSliceIndexMap(headerSlice)

	var records []record

	for {
		recordSlice, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(fmt.Sprintf("Error reading record slice with %v", err))
		}

		recordStruct := record{
			headerIndexMap: &headerIndexMap,
			values: recordSlice,
		}

		records = append(records, recordStruct)
	}

	return records
}