package model

type Todo struct {
	ID    			uint `gorm:"primaryKey" json:"id"`
	Title 			string `json:"title"`
	Status      string `json:"status"`
}
