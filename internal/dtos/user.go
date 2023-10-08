package dtos

type CreateUser struct {
	ExternalId string `json:"external_id" binding:"required"`
}