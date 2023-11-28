package helpers

import (
	"errors"
	"fmt"
	"bytes"
	"math/rand"
	"time"
)

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

func FindIntInSlice(intSlice []int, target int) bool {
	for _, integer := range intSlice {
		if integer == target {
			return true
		}
	}

	return false
}


func GetSliceIndexMap(slice []string) map[string]int {
	sliceIndexMap := make(map[string]int)
	for index, val := range slice {
		sliceIndexMap[val] = index
	}

	return sliceIndexMap
}

func RemoveBOM(bomString string) string {
	if bytes.Equal([]byte(bomString[:3]), []byte{0xEF, 0xBB, 0xBF}) {
		return bomString[3:]
	}

	return bomString
}

func GetKeySliceFromMap(stringMap map[string]bool) []string {
	keySlice := make([]string, len(stringMap))

	i := 0
	for key := range stringMap {
		keySlice[i] = key
		i++
	}

	return keySlice
}

func ShuffleSlice(slice []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}


