package auth

type RequestLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RequestLoginDummy struct {
	Username string `form:"userLogin" json:"userLogin" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
