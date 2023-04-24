package repository

import (
	"REST/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type reportRepository struct {
	db *sqlx.DB
}

func NewReportRepository(db *sqlx.DB) *reportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) CreateReport(report *model.Report) error {
	q := `
		INSERT INTO report(title, description)
		VALUES ($1, $2);`

	fmt.Println(report)
	_, err := r.db.Exec(q, &report.Title, &report.Description)

	return err

}

func (r *reportRepository) GetAllReport() ([]model.Report, error) {
	q := `
			SELECT * FROM report
`
	var output []model.Report

	err := r.db.Select(&output, q)

	if err != nil {
		return nil, err
	}

	return output, nil

}

func (r *reportRepository) GetByID(id int) (*model.Report, error) {
	q := `
			SELECT *
			FROM report
			WHERE reportid = $1;`
	var output model.Report
	err := r.db.Get(&output, q, id)

	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (r *reportRepository) DeleteReport(id int) error {
	q := `  DELETE from report
			WHERE reportid = $1;`

	_, err := r.db.Exec(q, id)

	return err

}

func (r *reportRepository) InitDB(query string) (interface{}, error) {

	result, err := r.db.Exec(query)
	if err != nil {
		log.Println("ошибка схемы запроса")
		log.Fatal(result)
		return result, err
	}

	return result, nil

}
