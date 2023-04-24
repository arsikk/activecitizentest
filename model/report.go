package model

type Report struct {
	ReportID    int     `db:"reportid" json:"reportid"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	UserID      *int    `db:"userid" json:"userID"`
	PhotoUrl    *string `db:"photourl" json:"photoUrl"`
}
