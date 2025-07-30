package config

type DB struct {
	Host string `db:"host"`
	Port int    `db:"port"`
	User string `db:"user"`
	Pass string `db:"pass"`
	Name string `db:"name"`
}
