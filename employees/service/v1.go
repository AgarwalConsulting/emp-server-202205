package service

import (
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
)

type svcV1 struct {
	repo repository.EmployeeRepository
}

func (s *svcV1) Index() ([]entities.Employee, error) {
	return s.repo.List()
}

func (s *svcV1) Create(newEmp entities.Employee) (*entities.Employee, error) {
	return s.repo.Create(newEmp)
}

func NewV1(repo repository.EmployeeRepository) EmployeeService {
	return &svcV1{repo}
}
