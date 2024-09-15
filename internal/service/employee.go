package service

import (
	"conducting-tenders/internal/entity"
	"conducting-tenders/internal/repo"
	"context"
)

type EmployeeService struct {
	employeeRepo repo.Employee
}

func NewEmployeeService(employeeRepo repo.Employee) *EmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
	}
}

func (b *EmployeeService) GetEmployees(ctx context.Context, input GetEmployeesInput) ([]entity.Employee, error) {
	employees, err := b.employeeRepo.GetEmployees(ctx, input.Limit)

	if err != nil {
		return []entity.Employee{}, err
	}

	return employees, nil
}
