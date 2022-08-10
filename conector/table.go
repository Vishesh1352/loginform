package conector

type LoginData struct {
	ID       uint   `json:"Id" gorm:"primary_key;auto_increment;not_null"`
	Name     string `json:"Name" gorm:"not_null,Unique"`
	Password string `json:"Password" gorm:"passwd,not_null"`
	Address  string `json:"Address" gorm:"not_null"`
	Image    string `json:"image"`
}
