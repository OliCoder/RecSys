package models

type Movie struct {
	MovieId   int64   `json:"movieid" gorm:"column:movieid"`
	Title     string  `json:"title"`
	Genre     string  `json:"genre"`
	AvgRating float32 `json:"avgrating" gorm:"column:avgrating"`
}
