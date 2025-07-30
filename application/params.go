package application

import (
	"flag"
)

type Params struct {
	DB    string
	Table string
	Clean bool
}

func NewParams() *Params {
	p := &Params{}
	flag.StringVar(&p.DB, "db", "", "database connection name")
	flag.StringVar(&p.Table, "table", "", "table pattern in regular expression syntax")
	flag.BoolVar(&p.Clean, "clean", false, "do not generate commented-out existing indexes")
	return p
}

func (p *Params) Load() {
	flag.Parse()
}

func (p *Params) Valid() bool {
	return p.DB != "" && p.Table != ""
}

func (p *Params) Help() {
	flag.PrintDefaults()
}
