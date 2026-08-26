package repository

import (
	"database/sql"
	"errors"

	"github.com/Ludimila-Araujo/labsupply-payment/internal/domain"
	"github.com/google/uuid"
)

type PostgresPaymentRepository struct {
	db *sql.DB
}

func NewPostgresPaymentRepository(db *sql.DB) *PostgresPaymentRepository {
	return &PostgresPaymentRepository{
		db: db,
	}
}

func (r *PostgresPaymentRepository) Create(
	payment *domain.Payment,
) error {

	_, err := r.db.Exec(`
		INSERT INTO payments (
			id,
			order_id,
			amount,
			status,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`,
		payment.ID,
		payment.OrderID,
		payment.Amount,
		payment.Status,
		payment.CreatedAt,
		payment.UpdatedAt,
	)

	return err
}

func (r *PostgresPaymentRepository) FindByID(
	id uuid.UUID,
) (*domain.Payment, error) {

	payment := &domain.Payment{}

	err := r.db.QueryRow(`
		SELECT
			id,
			order_id,
			amount,
			status,
			created_at,
			updated_at
		FROM payments
		WHERE id = $1
	`, id).Scan(
		&payment.ID,
		&payment.OrderID,
		&payment.Amount,
		&payment.Status,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrPaymentNotFound
		}

		return nil, err
	}

	return payment, nil
}

func (r *PostgresPaymentRepository) Update(
	payment *domain.Payment,
) error {

	_, err := r.db.Exec(`
		UPDATE payments
		SET
			status = $1,
			updated_at = $2
		WHERE id = $3
	`,
		payment.Status,
		payment.UpdatedAt,
		payment.ID,
	)

	return err
}
