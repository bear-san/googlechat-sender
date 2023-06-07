package main

import (
	"context"
	"github.com/bear-san/googlechat-sender/backend/cmd"
	"github.com/bear-san/googlechat-sender/backend/ent/postschedule"
	auth2 "github.com/bear-san/googlechat-sender/backend/internal/auth"
	"github.com/bear-san/googlechat-sender/backend/internal/db"
	"github.com/bear-san/googlechat-sender/backend/pkg/auth"
	"github.com/bear-san/googlechat-sender/backend/pkg/chat"
	"os"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		cmd.Start()
	} else if args[1] == "scheduler" {
		// 予約投稿の処理用
		ctx := context.Background()
		t := time.Now()

		schedules, err := db.Client.PostSchedule.Query().
			Where(postschedule.IsSentEQ(false)).
			All(ctx)

		if err != nil {
			panic(err)
		}

		for _, schedule := range schedules {
			if schedule.SendAt.Unix() <= t.Unix() {
				u, err := db.Client.SystemUser.Get(ctx, schedule.UID)
				if err != nil {
					continue
				}

				token, err := auth.GetGoogleCredential(ctx, u, auth2.GetOAuthClientInfo())
				if err != nil {
					continue
				}

				_, err = chat.Publish(ctx, token, schedule)
			}
		}
	}
}
