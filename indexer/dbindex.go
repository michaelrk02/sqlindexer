package indexer

type DBIndex struct {
	NonUnique  int    `db:"Non_unique"`
	KeyName    string `db:"Key_name"`
	SeqInIndex int    `db:"Seq_in_index"`
	ColumnName string `db:"Column_name"`
}

func GroupDBIndexes(table string, dbIndexes []DBIndex) []Index {
	idxMap := make(map[string]*Index)
	idxOrder := make(map[string]int)

	for _, dbIndex := range dbIndexes {
		if _, ok := idxMap[dbIndex.KeyName]; !ok {
			idxOrder[dbIndex.KeyName] = len(idxMap)
			idxMap[dbIndex.KeyName] = &Index{
				Name:   dbIndex.KeyName,
				Table:  table,
				Fields: []string{},
			}
		}

		idxMap[dbIndex.KeyName].Fields = append(idxMap[dbIndex.KeyName].Fields, dbIndex.ColumnName)
	}

	idxArr := make([]Index, len(idxMap))
	for key, value := range idxMap {
		idxArr[idxOrder[key]] = *value
	}

	return idxArr
}
