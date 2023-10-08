package dtos

type CreateSubscription struct {
	Type string `json:"type" binding:"required,oneof=push email"`
	Token string `json:"token" binding:"required"`
}

type UpdateSubscription struct {
	Type string `json:"type" binding:"required,oneof=push email"`
	Token string `json:"token" binding:"required"`
}