package domain

type Organization struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    Name      string `json:"name"`
    Document  string `json:"document"`
}
