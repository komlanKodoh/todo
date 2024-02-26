package controller

import (
	"gorm.io/gorm"
	"time"
)

type RequestTodoDto struct {
	Title       *string `json:"Title"`
	Description *string `json:"Description"`
	Completed   *bool   `json:"Completed"`
}

type AddTodoDto struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type PatchTodoDto struct {
	Title       *string `json:"Title"`
	Description *string `json:"Description"`
	Completed   *bool   `json:"Completed"`
}

type GetTodoDto struct {
	Id          uint           `json:"Id"`
	Title       string         `json:"Title"`
	Description string         `json:"Description"`
	CompletedAt gorm.DeletedAt `json:"CompletedAt"`
	CreatedAt   time.Time      `json:"CreatedAt"`
	UpdatedAt   time.Time      `json:"UpdatedAt"`
	DeletedAt   gorm.DeletedAt `json:"DeletedAt"`
}
