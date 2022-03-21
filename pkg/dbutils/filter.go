package dbutils

import (
	"fmt"
	"strings"
)

func AddStringFilter(str, connector, key, value string) string {
	if !strings.Contains(strings.ToLower(str), "where") {
		str = fmt.Sprintf("%s WHERE %s LIKE '%s'", str, key, "%"+value+"%")
		return str
	}

	str = fmt.Sprintf("%s %s %s LIKE '%s'", str, connector, key, "%"+value+"%")
	return str
}

func AddBigintFilter(str, connector, key string, value int64) string {
	if !strings.Contains(strings.ToLower(str), "where") {
		str = fmt.Sprintf("%s WHERE %s = '%d'", str, key, value)
		return str
	}

	str = fmt.Sprintf("%s %s %s = '%d'", str, connector, key, value)
	return str
}

func AddCustomFilter(str, connector, key, comparator, param string) string {
	if !strings.Contains(strings.ToLower(str), "where") {
		str = fmt.Sprintf("%s WHERE %s %s %s", str, key, comparator, param)
		return str
	}

	str = fmt.Sprintf("%s %s %s %s %s", str, connector, key, comparator, param)
	return str
}
