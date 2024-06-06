package domain

type User struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Phone     string `json:"phone"`
    Document  string `json:"document"`
}
