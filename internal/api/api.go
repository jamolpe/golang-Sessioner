package api

import (
	authorizationcore "authorization-service/internal/authorization-core"
	sessioncore "authorization-service/internal/session-core"
)

// New : define a new api
func New(authSrv authorizationcore.AuthServiceInterface, sessionSrv sessioncore.SessionInterface) *API {
	return &API{authSrv: authSrv, sessionSrv: sessionSrv}
}

type (
	// API : core endpoint
	API struct {
		authSrv    authorizationcore.AuthServiceInterface
		sessionSrv sessioncore.SessionInterface
	}
)
