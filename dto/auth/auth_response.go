package authdto

type LoginResponse struct {
	Id     int    `gorm:"type: int" json:"id"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Status string `gorm:"type: varchar(50)"  json:"status"`
	Gender string `gorm:"type: varchar(255)" json:"gender"`
	Phone  string `gorm:"type: varchar(255)" json:"phone"`
	Token  string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
	Id     int    `gorm:"type: int" json:"id"`
	Name   string `gorm:"type: varchar(255)" json:"name"`
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Gender string `gorm:"type: varchar(255)" json:"gender"`
	Phone  string `gorm:"type: varchar(255)" json:"phone"`
	Status string `gorm:"type: varchar(50)"  json:"status"`
}

type RegisterResponse struct {
	Id       int    `json:"id"`
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Gender   string `gorm:"type: varchar(255)" json:"gender"`
	Phone    string `gorm:"type: varchar(255)" json:"phone"`
	Status   string `gorm:"type: varchar(100)" json:"status"`
}
