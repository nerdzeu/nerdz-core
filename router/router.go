/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package router

import (
	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/oauth2"
	"github.com/nerdzeu/nerdz-api/oauth2/appauth"
	"github.com/nerdzeu/nerdz-api/rest"
	"github.com/nerdzeu/nerdz-api/stream"
	"net/http"
	"strings"
)

// Init configures the router and returns the *echo.Echo struct
// enableLog set to true enable echo middleware logger
func Init(enableLog bool) *echo.Echo {
	e := echo.New()
	if enableLog {
		e.Use(mw.Logger())
	}

	// Create the Authorization server for OAuth2
	authConfig := osin.NewServerConfig()
	// Configure the Authorization server
	authConfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	authConfig.AllowedAccessTypes = osin.AllowedAccessType{
		osin.AUTHORIZATION_CODE,
		osin.REFRESH_TOKEN,
		osin.PASSWORD,
		osin.CLIENT_CREDENTIALS,
	}
	authConfig.AllowGetAccessRequest = true
	authConfig.AllowClientSecretInParams = true

	// Create the storage for osin (where to save oauth infos)
	var authStorage nerdz.OAuth2Storage
	authorizationServer := osin.NewServer(authConfig, &authStorage)

	// Initialize oauth2 server implementation
	oauth2.Init(authorizationServer)

	// OAuth2 routes
	o := e.Group("/oauth2")
	o.Get("/authorize", oauth2.Authorize())
	o.Post("/authorize", oauth2.Authorize())
	o.Get("/token", oauth2.Token())
	o.Get("/info", oauth2.Info())
	o.Get("/app", oauth2.App())

	aa := o.Group("/appauth")
	aa.Use(Authorize())
	aa.Get("/code", appauth.Code())
	aa.Get("/token", appauth.Token())
	aa.Get("/password", appauth.Password())
	aa.Get("/client_credentials", appauth.ClientCredentials())
	aa.Get("/refresh", appauth.Refresh())
	aa.Get("/info", appauth.Info())

	// Content routes: requires application/user is authorized
	users := e.Group("/users")
	users.Use(Authorize())
	users.Get("/:id/posts", rest.UserPosts())
	users.Get("/:id/friends", rest.UserFriends())
	users.Get("/:id", rest.UserInfo())

	// Stream API
	s := e.Group("/stream")
	s.Use(Authorize())
	// notification for current logged in user
	s.Get("/notifications", stream.Notifications())

	return e
}

// Authorize is the authorization middleware for users.
// It checks the access_token in the Authorization header or the access_token query parameter
func Authorize() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			var accessToken string
			auth := c.Request().Header().Get("Authorization")
			if auth == "" {
				// Check if there's the parameter access_token in the URL
				// this makes the bearer authentication with websockets compatible with OAuth2
				accessToken = c.Query("access_token")
				if accessToken == "" {
					return c.String(http.StatusUnauthorized, "access_token required")
				}
			} else {
				if !strings.HasPrefix(auth, "Bearer ") {
					return echo.ErrUnauthorized
				}
				ss := strings.Split(auth, " ")
				if len(ss) != 2 {
					return echo.ErrUnauthorized
				}
				accessToken = ss[1]
			}

			accessData, err := (&nerdz.OAuth2Storage{}).LoadAccess(accessToken)
			if err != nil {
				return c.String(http.StatusUnauthorized, err.Error())
			}

			// store the Access Data into the context
			c.Set("accessData", accessData)

			// let next handler handle the context
			return next.Handle(c)
		})
	}
}