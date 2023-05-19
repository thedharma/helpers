package helpers


import (
	"encoding/csv"
	"fmt"
	"io"
)


type Record struct {
	headerIndexMap *map[string]int
	values []string
}



func (r *Record) Lookup(header string) string {
	headerIndex, found := (*r.headerIndexMap)[header]
	if !found {
		panic("header title not found")
	}
	return r.values[headerIndex]
}

func CSVReaderToRecords(csvReader *csv.Reader) []Record {
	
	headerSlice, err := csvReader.Read()
	
	if err != nil {
		panic(fmt.Sprintf("Error reading header with %v", err))
	}

	headerIndexMap := GetSliceIndexMap(headerSlice)

	var records []Record

	for {
		recordSlice, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(fmt.Sprintf("Error reading record slice with %v", err))
		}

		recordStruct := Record{
			headerIndexMap: &headerIndexMap,
			values: recordSlice,
		}

		records = append(records, recordStruct)
	}

	return records
}