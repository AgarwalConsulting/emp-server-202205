package repository_test

import (
	"sync"
	"testing"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
	"github.com/stretchr/testify/assert"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMem()

	initialEmps, err := sut.List()

	assert.Nil(t, err)
	assert.NotNil(t, initialEmps)

	initialEmpCount := len(initialEmps)

	assert.NotEqual(t, 0, initialEmpCount)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			newEmp := entities.Employee{Name: "G", Department: "LnD", ProjectID: 1001}

			sut.Create(newEmp)
		}()
	}

	wg.Wait()

	finalEmps, err := sut.List()

	assert.Nil(t, err)
	assert.NotNil(t, initialEmps)

	finalEmpCount := len(finalEmps)
	assert.NotEqual(t, 0, finalEmpCount)
	assert.Equal(t, 100, finalEmpCount-initialEmpCount)
}
