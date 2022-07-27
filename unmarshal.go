package gocsv

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func Unmarshal(data []byte, v any) error {
	reflectT := reflect.TypeOf(v).Elem().Elem()
	reflectV := reflect.ValueOf(v).Elem()
	structsLen := reflectT.NumField()
	var vts = make([]reflect.Value, 0)
	csvContents := strings.Split(string(data), "\n")

	if len(csvContents) > 1 {
		headers, contents := csvContents[0], csvContents[1:]
		var headersMap = make(map[string]int)
		var headersSlice []string
		if headers = strings.TrimSpace(headers); headers == "" {
			return errors.New("文件头无内容！")
		} else {
			headersSlice = strings.Split(headers, ",")
			for i, h := range headersSlice {
				headersMap[h] = i
			}
		}

		for _, line := range contents {
			if line = strings.TrimSpace(line); line != "" {
				vt := reflect.New(reflectT).Elem()
				vals := strings.Split(line, ",")
				for i := 0; i < structsLen; i++ {
					tag := reflectT.Field(i).Tag.Get("csv")
					value := vals[headersMap[tag]]
					switch reflectT.Field(i).Type.Kind().String() {
					case "int", "int8", "int16", "int32", "int64":
						intVal, err := strconv.Atoi(value)
						if err != nil {
							return err
						}
						vt.Field(i).SetInt(int64(intVal))
						break
					case "uint", "uint8", "uint16", "uint32", "uint64":
						uintVal, err := strconv.Atoi(value)
						if err != nil {
							return err
						}
						vt.Field(i).SetUint(uint64(uintVal))
						break
					case "float32", "float64":
						floatVal, err := strconv.ParseFloat(value, 64)
						if err != nil {
							return err
						}
						vt.Field(i).SetFloat(floatVal)
						break
					case "bool":
						if value == "true" {
							vt.Field(i).SetBool(true)
						} else {
							vt.Field(i).SetBool(false)
						}
						break
					case "string":
						vt.Field(i).SetString(value)
						break
					default:
						vt.Field(i).SetString(value)
					}
				}
				vts = append(vts, vt)
			}
		}
	}
	reflectV.Set(reflect.Append(reflectV, vts...))
	return nil
}
