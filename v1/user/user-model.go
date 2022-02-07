package user

import "time"

type General struct {
	// ID        int       `json:"id"`
	Username string `json:"username" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	// Status    bool      `json:"status" binding:"required"`

}

type CreateAtBy struct {
	CreatedBy string `json:"created_by" binding:"required"`
}

type UpdateAtBy struct {
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by" binding:"required"`
}

type CreateReq struct {
	General
}
