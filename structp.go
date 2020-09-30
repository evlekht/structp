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

func printField(field reflect.Value, fieldTypeOriginal reflect.Type, fieldName string, diveIntoPointers bool, indent int, lineEnding string) string {
	switch field.Type().Kind() {
	case reflect.Struct:
		nestedOutput := print(field.Interface(), diveIntoPointers, indent, true, lineEnding)
		return fmt.Sprintf("%s%s %v : %s,%s", getIndent(indent), fieldName, fieldTypeOriginal, nestedOutput, lineEnding)
	case reflect.Ptr:
		return fmt.Sprintf("%s%s %v : %v,%s", getIndent(indent), fieldName, fieldTypeOriginal, field.Pointer(), lineEnding)
	default:
		return fmt.Sprintf("%s%s %v : %v,%s", getIndent(indent), fieldName, fieldTypeOriginal, field.Interface(), lineEnding)
	}
}

func print(v interface{}, diveIntoPointers bool, indent int, omitType bool, lineEnding string) (output string) {
	defer func() {
		if recover() != nil {
			output += " <PANIC> "
		}
	}()

	valueOfV := reflect.Indirect(reflect.ValueOf(v))
	typeOfV := valueOfV.Type()

	indentStr := getIndent(indent)

	if typeOfV.Kind() != reflect.Struct {
		return
	}

	if omitType {
		output = fmt.Sprintf("{%s", lineEnding)
	} else {
		output = fmt.Sprintf("%s%v {%s", indentStr, typeOfV, lineEnding)
	}

	for i := 0; i < typeOfV.NumField(); i++ {
		field := valueOfV.Field(i)
		fieldType := field.Type()
		if diveIntoPointers && fieldType.Kind() == reflect.Ptr {
			field = reflect.Indirect(field)
		}
		output += printField(field, fieldType, typeOfV.Field(i).Name, diveIntoPointers, indent+1, lineEnding)
	}
	output += fmt.Sprintf("%s}", indentStr)
	return
}

// Print will separate lines with \n
func Print(v interface{}, diveIntoPointers bool) string {
	return print(v, diveIntoPointers, 0, false, "\n")
}

func PrintWithCustomLineEnding(v interface{}, diveIntoPointers bool, lineEnding string) string {
	return print(v, diveIntoPointers, 0, false, lineEnding)
}
