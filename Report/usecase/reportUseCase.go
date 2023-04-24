package usecase

import (
	"REST/model"
	"context"
	"time"
)

type ReportRepo interface {
	CreateReport(report *model.Report) error
	GetAllReport() ([]model.Report, error)
	GetByID(id int) (*model.Report, error)
	DeleteReport(id int) error
	InitDB(query string) (interface{}, error)
}

type ReportRedis interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, value interface{}, key string, ttl time.Duration) error
}

type reportUseCase struct {
	reportRepo  ReportRepo
	ReportRedis ReportRedis
}

func (r *reportUseCase) CreateReport(report *model.Report) error {
	return r.reportRepo.CreateReport(report)
}

func (r *reportUseCase) GetAllReport() ([]model.Report, error) {
	return r.reportRepo.GetAllReport()
}

func (r *reportUseCase) GetByID(id int) (*model.Report, error) {
	return r.reportRepo.GetByID(id)
}

func (r *reportUseCase) DeleteReport(id int) error {
	return r.reportRepo.DeleteReport(id)
}

func (r *reportUseCase) InitDB(query string) (interface{}, error) {
	return r.reportRepo.InitDB(query)
}

func (r *reportUseCase) Get(ctx context.Context, key string) (interface{}, error) {
	return r.ReportRedis.Get(ctx, key)

}

func (r *reportUseCase) Set(ctx context.Context, value interface{}, key string, ttl time.Duration) error {
	return r.ReportRedis.Set(ctx, value, key, ttl)

}
