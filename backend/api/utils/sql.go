package utils

import "fmt"

func ConvertToSQLWhereCondition(field string) string {
	if field == "" {
		return "IS NOT NULL"
	}

	return fmt.Sprintf(" = '%s'", field)
}
