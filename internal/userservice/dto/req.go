package dto

type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}
