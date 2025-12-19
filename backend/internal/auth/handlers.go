package auth

import "github.com/gin-gonic/gin"

type Handler struct{}

type registerReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func (h *Handler) Register(c *gin.Context) {
	// TODO: implement
}

func (h *Handler) Login(c *gin.Context) {
	// TODO: implement
}