package main

import (
	"fmt"
	"regexp"

	"github.com/michaelrk02/sqlindexer/application"
	"github.com/michaelrk02/sqlindexer/config"
	"github.com/michaelrk02/sqlindexer/indexer"
)

func main() {
	params := application.NewParams()
	params.Load()
	if !params.Valid() {
		params.Help()
		return
	}

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	dbCfg := cfg.DB[params.DB]

	db, err := application.Connect(&dbCfg)
	if err != nil {
		panic(err)
	}

	ixr := indexer.NewIndexer(db, cfg.Pattern)

	tables, err := db.GetTables()
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(params.Table)

	for _, table := range tables {
		if !re.MatchString(table) {
			continue
		}

		existing := make(map[string]bool)

		indexes, err := ixr.GetTableIndexes(table)
		if err != nil {
			panic(err)
		}

		for _, idx := range indexes {
			existing[idx.ID()] = true
		}

		indexesToCreate, err := ixr.GetTableIndexesToCreate(table)
		if err != nil {
			panic(err)
		}

		for _, idx := range indexesToCreate {
			if _, ok := existing[idx.ID()]; !ok {
				fmt.Printf("%s;\n", idx.SQL())
			} else if !params.Clean {
				fmt.Printf("-- %s;\n", idx.SQL())
			}
		}
	}
}
