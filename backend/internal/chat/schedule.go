package chat

import (
	"context"
	"encoding/json"
	"github.com/bear-san/googlechat-sender/backend/ent/postschedule"
	"github.com/bear-san/googlechat-sender/backend/internal/db"
	"github.com/bear-san/googlechat-sender/backend/pkg/auth"
	"github.com/bear-san/googlechat-sender/backend/pkg/chat"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func ScheduleList(req *gin.Context) {
	ctx := context.Background()
	u, err := auth.CheckStatus(ctx, req, os.Getenv("SECRET_BASE"))
	if err != nil {
		req.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":      "error",
				"description": "unauthorized",
			},
		)

		return
	}

	lst, err := db.Client.PostSchedule.Query().
		Where(postschedule.UIDEQ(u.ID)).
		All(ctx)

	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "failed to get your scheduled messages",
			},
		)

		return
	}

	req.JSON(
		http.StatusOK,
		lst,
	)
}

func Schedule(req *gin.Context) {
	ctx := context.Background()
	u, err := auth.CheckStatus(ctx, req, os.Getenv("SECRET_BASE"))
	if err != nil {
		req.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":      "error",
				"description": "unauthorized",
			},
		)

		return
	}

	var reqPayload Request
	reqTxt, err := io.ReadAll(req.Request.Body)
	if err == nil {
		err = json.Unmarshal(reqTxt, &reqPayload)
	}

	if err != nil {
		req.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":      "error",
				"description": "bad request",
			},
		)

		return
	}

	if reqPayload.Message == nil || reqPayload.Space == nil || reqPayload.SendAt == nil {
		req.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":      "error",
				"description": "some parameters are missing",
			},
		)

		return
	}

	sendAt := time.Unix(int64(*reqPayload.SendAt), 0)

	res, err := chat.Schedule(ctx, u, *reqPayload.Space, *reqPayload.Message, sendAt)
	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "failed to schedule message",
			},
		)

		return
	}

	req.JSON(
		http.StatusOK,
		res,
	)
}

func Reschedule(req *gin.Context) {
	ctx := context.Background()
	u, err := auth.CheckStatus(ctx, req, os.Getenv("SECRET_BASE"))
	if err != nil {
		req.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":      "error",
				"description": "unauthorized",
			},
		)

		return
	}

	var reqPayload Request
	reqTxt, err := io.ReadAll(req.Request.Body)
	if err == nil {
		err = json.Unmarshal(reqTxt, &reqPayload)
	}

	if err != nil {
		req.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":      "error",
				"description": "bad request",
			},
		)

		return
	}

	if reqPayload.Message == nil || reqPayload.Space == nil || reqPayload.SendAt == nil {
		req.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":      "error",
				"description": "some parameters are missing",
			},
		)

		return
	}

	sendAt := time.Unix(int64(*reqPayload.SendAt), 0)

	res, err := chat.ReSchedule(ctx, req.Param("sid"), u, *reqPayload.Space, *reqPayload.Message, sendAt)
	if err != nil {
		req.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":      "error",
				"description": "failed to schedule message",
			},
		)

		return
	}

	req.JSON(
		http.StatusOK,
		res,
	)
}

func UnSchedule(req *gin.Context) {
	ctx := context.Background()
	u, err := auth.CheckStatus(ctx, req, os.Getenv("SECRET_BASE"))

	if err != nil {
		req.JSON(
			http.StatusUnauthorized,
			gin.H{
				"status":      "error",
				"description": "unauthorized",
			},
		)

		return
	}

	err = chat.UnSchedule(ctx, u, req.Param("sid"))
	if err != nil {
		req.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":      "error",
				"description": "failed to unschedule post",
			},
		)

		return
	}

	req.Status(http.StatusNoContent)
}

type Request struct {
	Space   *chat.Space   `json:"space,omitempty"`
	Message *chat.Message `json:"message,omitempty"`
	SendAt  *int          `json:"send_at,omitempty"`
}
