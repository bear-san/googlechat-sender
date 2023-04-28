package chat

import (
	"context"
	"github.com/bear-san/googlechat-sender/backend/ent"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func GetDirectMessages(ctx context.Context, token *ent.GoogleApiKey, sheetId string) (*[]DirectMessage, error) {
	nToken := oauth2.Token{
		AccessToken:  token.AccessToken,
		TokenType:    "Bearer",
		RefreshToken: token.RefreshToken,
		Expiry:       token.ExpirationDate,
	}
	tokenSrc := oauth2.StaticTokenSource(&nToken)

	client := option.WithTokenSource(tokenSrc)
	sheetSvc, err := sheets.NewService(ctx, client)
	if err != nil {
		return nil, err
	}

	users, err := sheetSvc.Spreadsheets.Values.Get(sheetId, "シート1!A2:G").Do()
	if err != nil {
		return nil, err
	}

	lst := make([]DirectMessage, 0)
	for _, u := range users.Values {
		if len(u) < 7 {
			continue
		}

		info := DirectMessage{
			EmployeeNumber: u[0].(string),
			Email:          u[1].(string),
			GoogleUserId:   u[2].(string),
			DisplayName:    u[6].(string),
		}

		lst = append(lst, info)
	}

	return &lst, nil
}

type DirectMessage struct {
	EmployeeNumber string `json:"employee_number"`
	Email          string `json:"email"`
	GoogleUserId   string `json:"google_user_id"`
	DisplayName    string `json:"display_name"`
	Space          *Space `json:"space,omitempty"`
}
