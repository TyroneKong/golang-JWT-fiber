package models

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Username string `gorm:"size:255;not null;unique" json:"username"`
	Password []byte `gorm:"size:255;not null;" json:"password"`
}
