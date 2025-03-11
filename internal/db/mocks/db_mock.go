package mocks

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type DatabaseMock struct {
	mock.Mock
}

func (d *DatabaseMock) Exec(query string, args ...interface{}) (sql.Result, error) {
	mockArgs := d.Called(query, args)

	return mockArgs.Get(0).(sql.Result), mockArgs.Error(1)
}

func (d *DatabaseMock) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (d *DatabaseMock) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}

func (d *DatabaseMock) Ping() error {
	return nil
}

type SqlResultMock struct {
    mock.Mock
}

func (s *SqlResultMock) LastInsertId() (int64, error) {
    return 0, nil
}

func (s *SqlResultMock) RowsAffected() (int64, error) {
    return 0, nil
}


