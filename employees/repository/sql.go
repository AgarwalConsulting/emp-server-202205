package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"algogrit.com/emp-server/entities"
)

type sqlRepo struct {
	*sql.DB
}

func (repo *sqlRepo) List() ([]entities.Employee, error) {
	rows, err := repo.DB.Query("SELECT * FROM employees")

	if err != nil {
		log.Println("Unable to retrieve:", err)
		return nil, err
	}

	var emps []entities.Employee
	for rows.Next() {
		var emp entities.Employee

		err := rows.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.ProjectID)

		if err != nil {
			log.Println("Unable to scan:", err)
			return nil, err
		}

		emps = append(emps, emp)
	}

	return emps, nil
}

func (repo *sqlRepo) Create(newEmp entities.Employee) (*entities.Employee, error) {
	var empCount int
	rows, err := repo.DB.Query("SELECT count(*) FROM employees")

	if err != nil {
		return nil, err
	}

	rows.Next()
	err = rows.Scan(&empCount)

	if err != nil {
		return nil, err
	}

	newEmp.ID = empCount + 1

	// Inserting into the table
	_, err = repo.DB.Exec("INSERT INTO employees (id, name, department, project_id) VALUES($1, $2, $3, $4)", newEmp.ID, newEmp.Name, newEmp.Department, newEmp.ProjectID)

	if err != nil {
		log.Println("Unable to insert:", err)
		return nil, err
	}

	return &newEmp, nil
}

func NewSQL(connStr string) EmployeeRepository {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalln("unable to connect:", err)
	}

	return &sqlRepo{db}
}
