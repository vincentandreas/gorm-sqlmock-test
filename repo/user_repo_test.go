package repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm-sqlmock-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}

func TestFindUser_shouldFound(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	implObj := NewImplementation(db)
	users := sqlmock.NewRows([]string{"id", "full_name", "user_name", "password"}).
		AddRow(1, "user", "user", "passwd")

	expectedSQL := "SELECT (.+) FROM \"users\" WHERE id =(.+)"
	mock.ExpectQuery(expectedSQL).WillReturnRows(users)
	_, res := implObj.FindUserById(1, context.TODO())
	assert.Nil(t, res)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestFindUser_shouldNotFound(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	implObj := NewImplementation(db)
	users := sqlmock.NewRows([]string{"id", "full_name", "user_name", "password"})

	expectedSQL := "SELECT (.+) FROM \"users\" WHERE id =(.+)"
	mock.ExpectQuery(expectedSQL).WillReturnRows(users)
	_, res := implObj.FindUserById(2, context.TODO())
	assert.True(t, errors.Is(res, gorm.ErrRecordNotFound))
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestAddUser_shouldSuccess(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	implObj := NewImplementation(db)

	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	expectedSQL := "INSERT INTO \"users\" (.+) VALUES (.+)"
	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
	mock.ExpectCommit()
	var reqUser models.User
	implObj.SaveUser(reqUser, context.TODO())
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateUserPasswordById_shouldSuccess(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	implObj := NewImplementation(db)

	updUserSQL := "UPDATE \"users\" SET .+"
	mock.ExpectBegin()
	mock.ExpectExec(updUserSQL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	_, err := implObj.UpdateUserPasswordById(1, "newpass", context.TODO())
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteUserById_shouldSuccess(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	implObj := NewImplementation(db)

	delSQL := "DELETE FROM \"users\" WHERE id = .+"
	mock.ExpectBegin()
	mock.ExpectExec(delSQL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := implObj.DeleteUserById(1, context.TODO())
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}
