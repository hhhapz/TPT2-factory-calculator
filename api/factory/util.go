package factory

import (
	"fmt"
	"strconv"
)

func convertStr(name string, source interface{}) (string, error) {
	str, ok := source.(string)
	if !ok {
		return "", fmt.Errorf("could not convert %s property to string, found %v instead", name, source)
	}
	return str, nil
}

func convertInt(name string, source interface{}) (int, error) {
	num, ok := source.(float64)
	if !ok {
		return 0, fmt.Errorf("could not convert %s property to int, found %v instead", name, source)
	}
	return int(num), nil
}

func convertInt64(name string, source interface{}) (int64, error) {
	num, ok := source.(float64)
	if !ok {
		return 0, fmt.Errorf("could not convert %s property to int64, found %v instead", name, source)
	}
	return int64(num), nil
}

func convertStrArr(name string, source interface{}) ([]string, error) {
	var arr []string
	interfaceArr, ok := source.([]interface{})
	if !ok {
		return nil, fmt.Errorf("could not convert %s property to []string, found %v instead", name, source)
	}
	for i, item := range interfaceArr {
		str, err := convertStr(name+"["+strconv.Itoa(i)+"]", item)
		if err != nil {
			return nil, err
		}

		arr = append(arr, str)
	}

	return arr, nil
}

func convertStrMap(name string, source interface{}) (map[string]string, error) {
	strMap := make(map[string]string)
	interfaceMap, ok := source.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("could not convert %s property to map[string]string, found %v instead", name, source)
	}
	for k, v := range interfaceMap {
		str, err := convertStr("map["+k+"]", v)
		if err != nil {
			return nil, err
		}
		strMap[k] = str
	}

	return strMap, nil
}
