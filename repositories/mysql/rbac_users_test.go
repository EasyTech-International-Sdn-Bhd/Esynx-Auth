package mysql

import (
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
	"github.com/easytech-international-sdn-bhd/esynx-auth/test"
	"strings"
	"testing"
)

func TestRbacUsersRepository_CreateUser(t *testing.T) {
	session, config := test.NewTestAuthProvider()
	db := NewMySqlDb()
	err := db.Open(config.GetConnection(), config.GetLogger())
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	session.Db = db.Engine
	repo := NewRbacUsersRepository(&session)
	err = repo.CreateUser(models.CreateRbacUser{
		Username:        "UnitTestUserName",
		Password:        "UnitTestPassword",
		ClientCompany:   "UnitTestClientCompany",
		Metadata:        "{}",
		Server:          "UnitTestServer",
		BiDealer:        "",
		BiSubscriptions: "",
		BiState:         "",
		BiIndustry:      "",
		CreatedBy:       strings.ToUpper("unitTestUserName"),
	})
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}
