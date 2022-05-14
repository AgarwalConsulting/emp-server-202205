package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestCreateV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.New(mockSvc)

	reqBody := `{"name": "Fenil", "speciality": "SRE", "project": 2001}`

	req := httptest.NewRequest("POST", "/v1/employees", strings.NewReader(reqBody))

	resRec := httptest.NewRecorder()

	empInfo := entities.Employee{Name: "Fenil", Department: "SRE", ProjectID: 2001}

	expectedEmp := empInfo
	expectedEmp.ID = 1

	mockSvc.EXPECT().Create(empInfo).Return(&expectedEmp, nil)

	sut.ServeHTTP(resRec, req)

	resp := resRec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var createdEmp entities.Employee
	err := json.NewDecoder(resp.Body).Decode(&createdEmp)

	assert.Nil(t, err)

	assert.Equal(t, empInfo, createdEmp)
}
