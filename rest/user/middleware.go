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

package user

import (
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
	"net/http"
	"strconv"
)

// SetOther is the middleware that checks if the current logged user can see the required profile
// and if the required profile exist.
// On success sets the "other" = *User variable in the context
func SetOther() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			var id uint64
			var e error
			if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Invalid user identifier specified",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			var other *nerdz.User
			if other, e = nerdz.NewUser(id); e != nil {
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "User does not exists",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			me := c.Get("me").(*nerdz.User)
			if !me.CanSee(other) {
				return c.JSON(http.StatusUnauthorized, &rest.Response{
					HumanMessage: "You can't see the required profile",
					Message:      "You can't see the required profile",
					Status:       http.StatusUnauthorized,
					Success:      false,
				})
			}

			// store the other User into the context
			c.Set("other", other)
			// pass context to the next handler
			return next.Handle(c)
		})
	}
}

// SetPost is the middleware that checks if the required post, on the user board, exists.
// If it exists, set the "post" = *UserPost in the current context
func SetPost() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			var e error
			var pid uint64

			if pid, e = strconv.ParseUint(c.Param("pid"), 10, 64); e != nil {
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Invalid post identifier specified",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			otherID := c.Get("other").(*nerdz.User).Counter
			var post *nerdz.UserPost

			if post, e = nerdz.NewUserPostWhere(&nerdz.UserPost{nerdz.Post{To: otherID, Pid: pid}}); e != nil {
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Required post does not exists",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			c.Set("post", post)
			return next.Handle(c)
		})
	}
}
