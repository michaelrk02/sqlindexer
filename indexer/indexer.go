package indexer

import (
	"fmt"
	"regexp"

	"github.com/michaelrk02/sqlindexer/application"
	"github.com/michaelrk02/sqlindexer/config"
)

type Indexer struct {
	db       *application.DB
	patterns []config.Pattern
}

func NewIndexer(db *application.DB, patterns []config.Pattern) *Indexer {
	return &Indexer{
		db:       db,
		patterns: patterns,
	}
}

func (ixr *Indexer) GetTableIndexes(table string) ([]Index, error) {
	var dbIndexes []DBIndex

	err := ixr.db.Select(&dbIndexes, fmt.Sprintf("SHOW INDEXES FROM `%s`", table))
	if err != nil {
		return nil, err
	}

	return GroupDBIndexes(table, dbIndexes), nil
}

func (ixr *Indexer) GetTableIndexesToCreate(table string) ([]Index, error) {
	columns, err := ixr.db.GetTableColumns(table)
	if err != nil {
		return nil, err
	}

	groupMap := make(map[string]*Index)
	groupIdx := make(map[string]int)

	for _, pat := range ixr.patterns {
		for _, tup := range pat.Tuple {
			re := regexp.MustCompile(tup)

			for _, col := range columns {
				match := re.FindStringSubmatch(col)
				if match != nil {
					keyName := GetKeyName(table, match[1])
					groupName := fmt.Sprintf("%s-%s", pat.ID, keyName)

					if _, ok := groupMap[groupName]; !ok {
						groupIdx[groupName] = len(groupMap)
						groupMap[groupName] = &Index{
							Name:      keyName,
							GroupName: groupName,
							Table:     table,
							Fields:    []string{},
						}
					}
					groupMap[groupName].Fields = append(groupMap[groupName].Fields, col)
				}
			}
		}
	}

	groupList := make([]Index, len(groupMap))
	for _, group := range groupMap {
		groupList[groupIdx[group.GroupName]] = *groupMap[group.GroupName]
	}
	return groupList, nil
}
