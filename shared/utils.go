package shared

func GetStringFromRow(row []interface{}, position int, empty string) string {
	if row == nil {
		return empty
	} else {
		if len(row) > position {
			if row[position] == nil {
				return empty
			} else {
				return row[position].(string)
			}
		} else {
			return empty
		}
	}
}
