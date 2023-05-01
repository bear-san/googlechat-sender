package chat

import (
	"context"
	"fmt"
	"github.com/bear-san/googlechat-sender/backend/ent"
	"github.com/bear-san/googlechat-sender/backend/internal/db"
	"github.com/google/uuid"
	"time"
)

func Schedule(ctx context.Context, u *ent.SystemUser, space Space, msg Message, sendAt time.Time) (*ent.PostSchedule, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return db.Client.PostSchedule.Create().
		SetID(id).
		SetUID(u.ID).
		SetTarget(*space.Name).
		SetDisplayName(*space.DisplayName).
		SetText(*msg.Text).
		SetIsSent(false).
		SetSendAt(sendAt).
		Save(ctx)
}

func ReSchedule(ctx context.Context, id string, u *ent.SystemUser, space Space, msg Message, sendAt time.Time) (*ent.PostSchedule, error) {
	rid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	s, err := db.Client.PostSchedule.Get(ctx, rid)
	if err != nil || s.UID != u.ID {
		return nil, fmt.Errorf("not found")
	}

	return db.Client.PostSchedule.UpdateOneID(rid).
		SetTarget(*space.Name).
		SetText(*msg.Text).
		SetSendAt(sendAt).
		Save(ctx)
}

func UnSchedule(ctx context.Context, u *ent.SystemUser, id string) error {
	rid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	s, err := db.Client.PostSchedule.Get(ctx, rid)
	if err != nil || s.UID != u.ID {
		return fmt.Errorf("not found")
	}

	return db.Client.PostSchedule.DeleteOneID(rid).Exec(ctx)
}

func Publish(ctx context.Context, cred *ent.GoogleApiKey, s *ent.PostSchedule) (*ent.PostSchedule, error) {
	space := Space{
		Name: &s.Target,
	}

	msg := Message{
		Text: &s.Text,
	}

	_, err := space.Post(cred, msg)
	if err != nil {
		return nil, err
	}

	return db.Client.PostSchedule.UpdateOneID(s.ID).
		SetIsSent(true).
		Save(ctx)
}
