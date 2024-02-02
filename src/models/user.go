package models

import "database/sql"

type User struct {
	ID         uint   `gorm:"primaryKey"`
	FirstName  string `gorm:"index:,sort:desc,"`
	LastName   string `gorm:"index:,sort:desc,"`
	Email      string `gorm:"unique"`
	CreatedAt  int    `gorm:"autoCreateTime"`
	Password   string
	About      sql.NullString
	Title      sql.NullString
	ProfileUrl sql.NullString
}
