package domain

import "time"

type SupAdmins struct {
	ID        uint   `gorm:"primaryKey;unique;not null"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Mobile    string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdateAt  time.Time
}