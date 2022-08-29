package gosnow

type TableResponse Response

// First returns the first record in the map
func (R TableResponse) First() (TableEntry, error) {
	content, _, err := R.All()
	if err != nil {
		//err = fmt.Errorf("could not retrieve first record because of upstream error")
		return TableEntry{}, err
	}

	if len(content) != 0 {
		logger.Println(content[0])
		return content[0], nil
	} else {
		return TableEntry{}, nil
	}
}

// All returns all found serviceNow records in a map slice
func (R TableResponse) All() ([]TableEntry, int, error) {

	entries, count, err := Response(R)._get_buffered_response()

	new_entries := []TableEntry{}

	for _, entry := range entries {
		new_entries = append(new_entries, TableEntry(entry))
	}

	return new_entries, count, err
}
