package repository

import (
	"sync"

	"algogrit.com/emp-server/entities"
)

type inMem struct {
	employees []entities.Employee
	mut       sync.RWMutex
}

func (repo *inMem) List() ([]entities.Employee, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()

	return repo.employees, nil
}

func (repo *inMem) Create(newEmployee entities.Employee) (*entities.Employee, error) {
	repo.mut.Lock()
	defer repo.mut.Unlock()
	newEmployee.ID = len(repo.employees) + 1
	repo.employees = append(repo.employees, newEmployee)

	return &newEmployee, nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Shoba", "SRE", 1002},
		{3, "Naveen", "Cloud", 10010},
	}

	return &inMem{employees: employees}
}
