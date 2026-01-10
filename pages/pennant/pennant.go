package pennant

import (
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/pennant/handlers"
)

func GetHandlers() []*common.Handler {
	return []*common.Handler{
		common.NewHandler("/pennant", "GET", handlers.WeekendCompetitionHandler),
		common.NewHandler("/pennant/:competition", "GET", handlers.CompetitionHandler),
		common.NewHandler("/pennant/authenticate", "POST", handlers.AuthenticationHandler),
		common.NewHandler("/pennant/lock", "GET", handlers.LockAuthenticationHandler),
	}
}