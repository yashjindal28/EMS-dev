package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllPayrollDataNoError(t *testing.T) {
	cursor, err := FindAllPayrollData()
	assert.Equal(t, nil, err, "Nil error is expected")
	assert.NotNil(t, cursor, "Elements are expected")
}

func TestGetPayrollInfoOfEmployeeByIDNoError(t *testing.T) {
	result := GetPayrollInfoOfEmployeeByID("EXE1")
	assert.NotNil(t, result, "Record should exist")
}
