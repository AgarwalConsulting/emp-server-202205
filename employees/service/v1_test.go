package service_test

import (
	"testing"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockEmployeeRepository(ctrl)

	sut := service.NewV1(mockRepo)

	expectedEmps := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	mockRepo.EXPECT().List().Return(expectedEmps, nil)

	emps, err := sut.Index()

	assert.Nil(t, err)
	assert.NotNil(t, emps)

	assert.Equal(t, expectedEmps, emps)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockEmployeeRepository(ctrl)

	sut := service.NewV1(mockRepo)

	newEmployee := entities.Employee{
		Name:       "Nageswari",
		Department: "Cloud",
		ProjectID:  1002,
	}

	expectedEmp := newEmployee
	expectedEmp.ID = 2

	mockRepo.EXPECT().Create(newEmployee).Return(&expectedEmp, nil)

	createdEmp, err := sut.Create(newEmployee)

	assert.Nil(t, err)
	assert.NotNil(t, createdEmp)

	assert.Equal(t, &expectedEmp, createdEmp)
}
