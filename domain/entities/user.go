package entities

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `gorm:"type:varchar(50)" json:"name"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}
