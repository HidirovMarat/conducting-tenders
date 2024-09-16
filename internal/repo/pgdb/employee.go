package pgdb

import (
	"conducting-tenders/internal/entity"
	"conducting-tenders/internal/repo/repoerrs"
	"conducting-tenders/pkg/postgres"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type EmployeeRepo struct {
	*postgres.Postgres
}

func NewEmployeeRepo(pg *postgres.Postgres) *EmployeeRepo {
	return &EmployeeRepo{pg}
}

func (r *EmployeeRepo) GetEmployeeByUsername(ctx context.Context, username string) (entity.Employee, error) {
	sql, args, _ := r.Builder.
		Select("id", "username", "first_name", "last_name").
		From("employee").
		Where("username = ?", username).
		ToSql()

	var employee entity.Employee
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&employee.Id,
		&employee.Username,
		&employee.First_name,
		&employee.Last_name,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Employee{}, repoerrs.ErrNotFound
		}
		return entity.Employee{}, fmt.Errorf("EmployeeRepo.GetEmployeeByUsername - r.Pool.QueryRow: %v", err)
	}

	return employee, nil
}

func (r *EmployeeRepo) GetEmployeeById(ctx context.Context, employeeId uuid.UUID) (entity.Employee, error) {
	sql, args, _ := r.Builder.
		Select("id", "username", "first_name", "last_name", "created_at", "update_at").
		From("employee").
		Where("id = ?", employeeId).
		ToSql()

	var employee entity.Employee
	err := r.Pool.QueryRow(ctx, sql, args...).Scan(
		&employee.Id,
		&employee.Username,
		&employee.First_name,
		&employee.Last_name,
		&employee.Created_at,
		&employee.Updated_at,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Employee{}, repoerrs.ErrNotFound
		}
		return entity.Employee{}, fmt.Errorf("EmployeeRepo.GetEmployeeById - r.Pool.QueryRow: %v", err)
	}

	return employee, nil
}

func (r *EmployeeRepo) GetEmployees(ctx context.Context, limit int) ([]entity.Employee, error) {
	req := r.Builder.
		Select("*").
		From("employee")

	if limit > 0 {
		req.Limit(uint64(limit))
	}

	sql, args, _ := req.ToSql()

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return []entity.Employee{}, fmt.Errorf("EmployeeRepo.GetEmployees - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	var employees []entity.Employee
	for rows.Next() {
		var employee entity.Employee
		err := rows.Scan(
			&employee.Id,
			&employee.Username,
			&employee.First_name,
			&employee.Last_name,
			&employee.Created_at,
			&employee.Updated_at,
		)
		if err != nil {
			return nil, fmt.Errorf("EmployeeRepo.GetEmployees - rows.Scan: %v", err)
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
