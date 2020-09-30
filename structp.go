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

func print(v interface{}, diveIntoPointers bool, indent int, omitType bool) (output string) {
	defer func() {
		if recover() != nil {
			output += " <PANIC> "
		}
	}()

	valueOfV := reflect.Indirect(reflect.ValueOf(v))
	typeOfV := valueOfV.Type()

	if typeOfV.Kind() != reflect.Struct {
		return
	}

	indentStr := getIndent(indent)
	innerIndentStr := indentStr + indentStrConst

	if omitType {
		output = fmt.Sprintf("{\n")
	} else {
		output = fmt.Sprintf("%s%v {\n", indentStr, typeOfV)
	}

	for i := 0; i < typeOfV.NumField(); i++ {
		field := valueOfV.Field(i)
		fieldType := field.Type()
	L:
		for {
			fieldTypeIndirect := field.Type()
			switch fieldTypeIndirect.Kind() {
			case reflect.Struct:
				nestedOutput := print(field.Interface(), diveIntoPointers, indent+1, true)
				output += fmt.Sprintf("%s%s %s : %s,\n", innerIndentStr, typeOfV.Field(i).Name, fieldType, nestedOutput)
				break L
			case reflect.Ptr:
				if diveIntoPointers {
					field = reflect.Indirect(field)
				} else {
					output += fmt.Sprintf("%s%s %s : %v,\n", innerIndentStr, typeOfV.Field(i).Name, fieldType, field.Pointer())
					break L
				}
			default:
				output += fmt.Sprintf("%s%s %s : %v,\n", innerIndentStr, typeOfV.Field(i).Name, fieldType, field.Interface())
				break L
			}
		}
	}
	output += fmt.Sprintf("%s}", indentStr)
	return
}

func Print(v interface{}, diveIntoPointers bool) string {
	return print(v, diveIntoPointers, 0, false)
}
