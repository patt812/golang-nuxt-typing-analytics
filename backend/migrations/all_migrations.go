package migrations

import "github.com/go-gormigrate/gormigrate/v2"

func GetAllMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		InsertKanaPattern(),
	}
}
