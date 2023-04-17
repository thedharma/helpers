package helpers

import (
	"errors"
	"fmt"
)

// Splits string into slice of strings at seperator which is not in quotes.
// Seperator must be able to be represented by one byte. Otherwise will panic with index error.
func SplitUnquoted(s, sep string) []string {
	result := []string{}
	byteSep := []byte(sep)[0]
	var begin int
	var inQuotes bool

	for i := 0; i < len(s); i++ {
		if s[i] == byteSep && !inQuotes {
			result = append(result, s[begin : i])
			begin = i + 1
		} else if s[i] == '"' {
			if !inQuotes {
				inQuotes = true
			} else {
				inQuotes = false
				}
		}
		
		
	}
	return append(result, s[begin :])
}

func StringMapFrom2Slices(keys, values []string) (map[string]string, error) {
	var stringMap = map[string]string{}
	keyLength, valueLength := len(keys), len(values)

	if len(keys) != len(values) {
		errMsg := fmt.Sprintf("difference in amount of keys and values: %v != %v", keyLength, valueLength)
		return stringMap, errors.New(errMsg)
	}

	for i := 0; i < keyLength; i++ {
		stringMap[keys[i]] = values[i]  
	}

	return stringMap, nil
}

