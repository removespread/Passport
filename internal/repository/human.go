package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"passport/internal/models"

	_ "github.com/lib/pq"
)

type HumanRepository struct {
	db *sql.DB
}

// HumanRepository интерфейс для работы с данными человека
type HumanRepositoryInterface interface {
	CreateHuman(human *models.Human) error
	UpdateHuman(human *models.Human) error
	DeleteHuman(human *models.Human) error
	GetHuman(human *models.Human) error
	GetAllHumans() ([]*models.Human, error)
	GetHumanBySerialNumber(serialNumber string) (*models.Human, error)
	InitDB() error
}

// NewHumanRepository создает новый репозиторий для работы с данными человека
func NewHumanRepository(host, port, user, password, dbname string) *HumanRepository {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	return &HumanRepository{db: db}
}

// CreateHuman создает нового человека в базе данных
func (r *HumanRepository) CreateHuman(human *models.Human) error {
	query := "INSERT INTO humans (first_name, last_name, surname, dob, serial_number, address, code_structure) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := r.db.Exec(query, human.FirstName, human.LastName, human.Surname, human.DOB, human.SerialNumber, human.Address, human.CodeStructure)
	if err != nil {
		return err
	}

	return nil
}

// GetHuman получает человека по ID
func (r *HumanRepository) GetHuman(id int) (*models.Human, error) {
	query := "SELECT * FROM humans WHERE id = $1"
	var human models.Human
	err := r.db.QueryRow(query, id).Scan(&human.ID, &human.FirstName, &human.LastName, &human.Surname, &human.SerialNumber, &human.Address, &human.CodeStructure)
	if err != nil {
		return nil, err
	}

	return &human, nil
}

// UpdateHuman обновляет данные человека в базе данных
func (r *HumanRepository) UpdateHuman(human *models.Human) error {
	query := "UPDATE humans SET first_name = $1, last_name = $2, surname = $3, dob = $4, serial_number = $5, address = $6, code_structure = $7 WHERE id = $8"
	_, err := r.db.Exec(query, human.FirstName, human.LastName, human.Surname, human.DOB, human.SerialNumber, human.Address, human.CodeStructure, human.ID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteHuman удаляет человека из базы данных
func (r *HumanRepository) DeleteHuman(human *models.Human) error {
	query := "DELETE FROM humans WHERE id = $1"
	_, err := r.db.Exec(query, human.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetAllHumans получает всех людей из базы данных
func (r *HumanRepository) GetAllHumans() ([]*models.Human, error) {
	query := "SELECT * FROM humans"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var humans []*models.Human
	for rows.Next() {
		var human models.Human
		err := rows.Scan(&human.ID, &human.FirstName, &human.LastName, &human.Surname, &human.SerialNumber, &human.Address, &human.CodeStructure)
		if err != nil {
			return nil, err
		}
		humans = append(humans, &human)
	}

	return humans, nil
}

// GetHumanBySerialNumber получает человека по серии паспорта
func (r *HumanRepository) GetHumanBySerialNumber(serialNumber string) (*models.Human, error) {
	query := "SELECT * FROM humans WHERE serial_number = $1"
	var human models.Human
	err := r.db.QueryRow(query, serialNumber).Scan(&human.ID, &human.FirstName, &human.LastName, &human.Surname, &human.SerialNumber, &human.Address, &human.CodeStructure)
	if err != nil {
		return nil, err
	}

	return &human, nil
}

// InitDB инициализирует базу данных и применяет миграции
func (r *HumanRepository) InitDB() error {
	// Читаем файл миграции
	migrationSQL, err := os.ReadFile("internal/migrations/001_init.sql")
	if err != nil {
		return fmt.Errorf("error reading migration file: %v", err)
	}

	// Выполняем миграцию
	_, err = r.db.Exec(string(migrationSQL))
	if err != nil {
		return fmt.Errorf("error executing migration: %v", err)
	}

	return nil
}
