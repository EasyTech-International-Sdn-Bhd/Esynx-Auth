package sql

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
	_ "github.com/go-sql-driver/mysql"
	"src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"
)

func DefineSchema(db *xorm.Engine) error {
	m := xormigrate.New(db, migrations())
	return m.Migrate()
}

func migrations() []*xormigrate.Migration {
	var schemas []*xormigrate.Migration
	for i, schema := range defaults() {
		schemas = append(schemas, &xormigrate.Migration{
			ID: fmt.Sprintf("define_schema_%d", i),
			Migrate: func(tx *xorm.Engine) error {
				return tx.Sync(schema)
			},
			Rollback: func(db *xorm.Engine) error {
				return nil
			},
		})
	}
	return schemas
}

func defaults() []interface{} {
	return []interface{}{
		&entities.RbacPermissions{},
		&entities.RbacRoles{},
		&entities.RbacUsers{},
		&entities.RbacUserRoles{},
		&entities.RbacRolesPermissions{},
		&entities.RbacTokens{},
	}
}
