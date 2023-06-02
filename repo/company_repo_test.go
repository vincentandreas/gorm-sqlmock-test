package repo

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImplementation_FindCompanyById(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()

	implObj := NewImplementation(db)
	companies := sqlmock.NewRows([]string{"id", "dummy_column"}).
		AddRow(1, "Abc lte")

	companySQL := "SELECT (.+) FROM \"companies\" JOIN company_scales ON companies.company_scale_id = company_scales.id WHERE id =.+"
	mock.ExpectQuery(companySQL).WillReturnRows(companies)
	_, res := implObj.FindCompanyById(1, context.TODO())
	assert.Nil(t, res)
	assert.Nil(t, mock.ExpectationsWereMet())

}

func TestImplementation_DeleteCompanyById(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	implObj := NewImplementation(db)

	delSQL := "UPDATE \"companies\" SET \"deleted_at\"=.+ WHERE id =.+ AND \"companies\".\"deleted_at\" IS NULL"
	mock.ExpectBegin()
	mock.ExpectExec(delSQL).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := implObj.DeleteCompanyById(1, context.TODO())
	assert.Nil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}
