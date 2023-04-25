package models

func GetParamMarks(length int) string {
	sql := ""
	for i := 0; i < length; i++ {
		sql += "?,"
	}
	sql += "?"
	return sql
}