package auth

import (
	"github.com/bear-san/googlechat-sender/backend/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(req *gin.Context) {
	redirectUrl, err := auth.CreateLoginUrl()
	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "internal server error",
			},
		)

		return
	}

	req.Redirect(
		http.StatusFound,
		*redirectUrl,
	)
}
