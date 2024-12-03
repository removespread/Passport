package models

type Human struct {
	ID            int    `yaml:"id" gorm:"primary_key"`          // ID человека
	FirstName     string `yaml:"first_name" gorm:"not null"`     // Имя
	LastName      string `yaml:"last_name" gorm:"not null"`      // Фамилия
	Surname       string `yaml:"surname" gorm:"not null"`        // Отчество
	DOB           string `yaml:"dob" gorm:"not null"`            // Дата рождения
	SerialNumber  string `yaml:"serial_number" gorm:"not null"`  // Серия паспорта
	Address       string `yaml:"address" gorm:"not null"`        // Адрес
	CodeStructure string `yaml:"code_structure" gorm:"not null"` // Код выданной структуры МВД
}
