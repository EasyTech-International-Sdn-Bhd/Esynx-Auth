package sql

import (
	"database/sql"
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	migrate "github.com/easytech-international-sdn-bhd/esynx-auth/migrate/sql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type SqlDb struct {
	Engine *xorm.Engine
}

func NewSqlDb() *SqlDb {
	return &SqlDb{}
}

func (m *SqlDb) Open(conn string, logger contracts.IDatabaseLogger) (err error) {
	m.Engine, err = xorm.NewEngine("postgres", conn, func(db *sql.DB) error {
		db.SetMaxOpenConns(1)
		db.SetMaxIdleConns(1)
		err := db.Ping()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	if logger != nil {
		m.Engine.SetLogger(logger)
	}
	m.Engine.ShowSQL(true)
	m.Engine.SetLogLevel(0)
	return nil
}

func (m *SqlDb) DefineSchema() error {
	return migrate.DefineSchema(m.Engine)
}

func (m *SqlDb) Close() error {
	err := m.Engine.Close()
	if err != nil {
		return err
	}
	return nil
}
