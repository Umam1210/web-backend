package authdto

type LoginResponse struct {
	ID      int    `json:"id"`
	Name    string `gorm:"type: varchar(255)" json:"name"`
	Email   string `gorm:"type: varchar(255)" json:"email"`
	Token   string `gorm:"type: varchar(255)" json:"token"`
	Address string `gorm:"type: varchar(255)" json:"address"`
	Phone   string `gorm:"type: varchar(255)" json:"phone"`
}
type CheckAuthResponse struct {
	Id    int    `gorm:"type: int" json:"id"`
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
}
