package structp

import (
	"fmt"
	"reflect"
)

const indentStrConst = "  "

func getIndent(indent int) string {
	output := ""
	for j := 0; j < indent; j++ {
		output += indentStrConst
	}
	return output
}

func print(v interface{}, indent int, omitType bool) string {
	indentStr := getIndent(indent)
	innerIndentStr := indentStr + indentStrConst
	valueOfV := reflect.ValueOf(v).Elem()
	typeOfV := valueOfV.Type()
	var output string
	if omitType {
		output = fmt.Sprintf("{\n")
	} else {
		output = fmt.Sprintf("%s%v {\n", indentStr, typeOfV)
	}
	for i := 0; i < typeOfV.NumField(); i++ {
		field := valueOfV.Field(i)
		if field.Type().Kind() == reflect.Struct {
			nestedOutput := print(field.Addr().Interface(), indent+1, true)
			output += fmt.Sprintf("%s%s %s : %s,\n", innerIndentStr, typeOfV.Field(i).Name, field.Type(), nestedOutput)
		} else {
			output += fmt.Sprintf("%s%s %s : %v,\n", innerIndentStr, typeOfV.Field(i).Name, field.Type(), field.Interface())
		}
	}
	output += fmt.Sprintf("%s}", indentStr)
	return output
}

func Print(v interface{}) string {
	return print(v, 0, false)
}
