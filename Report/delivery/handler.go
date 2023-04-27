package delivery

import (
	"REST/Report/usecase"
	"REST/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type handler struct {
	useCasePostgres usecase.ReportRepo
	useCaseRedis    usecase.ReportRedis
}

func NewHandler(useCase usecase.ReportRepo, useRedis usecase.ReportRedis) *handler {
	return &handler{useCasePostgres: useCase,
		useCaseRedis: useRedis}
}

func (h *handler) CreateReport(c *gin.Context) {
	type AddReportRequestBody struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	inp := AddReportRequestBody{}
	ctx := context.Background()

	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(inp)

	var report model.Report

	report.Title = inp.Title
	report.Description = inp.Description

	err := h.useCasePostgres.CreateReport(&report)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	err = h.useCaseRedis.Set(ctx, &report, strconv.Itoa(report.ReportID), 0)
	if err != nil {
		log.Printf("redis doesn't set cache")
	}

	c.JSON(http.StatusOK, "created")

}

func (h *handler) GetAllReport(c *gin.Context) {
	result, err := h.useCasePostgres.GetAllReport()

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, &result)

}

func (h *handler) GetByID(c *gin.Context) {
	inp := model.Report{}
	ctx := context.Background()

	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := h.useCaseRedis.Get(ctx, strconv.Itoa(inp.ReportID))
	if err == nil {
		log.Printf(" found in redis")
		c.JSON(http.StatusOK, result)
		return
	}

	result, err = h.useCasePostgres.GetByID(inp.ReportID)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, result)

}

func (h *handler) DeleteReport(c *gin.Context) {
	inp := model.Report{}
	ctx := context.Background()

	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.useCaseRedis.Del(ctx, strconv.Itoa(inp.ReportID))

	if err != nil {
		log.Printf("redis deleting error")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = h.useCasePostgres.DeleteReport(inp.ReportID)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, "deleted succesfully")
}
