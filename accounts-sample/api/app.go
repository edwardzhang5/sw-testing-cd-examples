package api

import (
	"github.com/drbyronw/accounts/db"
	"github.com/drbyronw/accounts/service"
)

// WebApp contains application components, structs, and interfaces
type WebApp struct {
	DB       *db.FSRepo
	Accounts service.AccountsService
}
