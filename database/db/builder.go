package db

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"

	"github.com/goravel/framework/support/str"
)

var NameMapper = func(s string) string {
	if s == "ID" {
		return "id"
	}

	return str.Of(s).Snake().String()
}

type Builder struct {
	*sqlx.DB
	gormDB *gorm.DB
}

func NewBuilder(gormDB *gorm.DB, driver string) (*Builder, error) {
	db, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	dbx := sqlx.NewDb(db, driver)
	dbx.MapperFunc(NameMapper)

	return &Builder{
		DB:     dbx,
		gormDB: gormDB,
	}, nil
}

func (r *Builder) Explain(sql string, args ...any) string {
	return r.gormDB.Explain(sql, args...)
}

type TxBuilder struct {
	*sqlx.Tx
	gormDB *gorm.DB
}

func NewTxBuilder(gormDB *gorm.DB, driver string) (*TxBuilder, error) {
	db, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	dbx := sqlx.NewDb(db, driver)
	dbx.MapperFunc(NameMapper)
	tx, err := dbx.Beginx()
	if err != nil {
		return nil, err
	}

	return &TxBuilder{
		Tx:     tx,
		gormDB: gormDB,
	}, nil
}

func (r *TxBuilder) Explain(sql string, args ...any) string {
	return r.gormDB.Explain(sql, args...)
}
