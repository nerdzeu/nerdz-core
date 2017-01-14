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

package rest

import (
	"github.com/nerdzeu/nerdz-api/nerdz"
)

// Response represents the response format of the API
// swagger:response apiResponse
type Response struct {
	// The API response data
	Data interface{} `json:"data"`
	// The API generated message
	Message string `json:"message"`
	// The human generated message, easy to understand
	HumanMessage string `json:"humanMessage"`
	// Status Code of the request
	Status uint `json:"status"`
	// Success indicates if the requested succeded
	Success bool `json:"success"`
}

// NewMessage represents a new message from the current user
// swagger:response message
type NewMessage struct {
	Message string `json:"message"`
	Lang    string `json:"lang, omitempty"`
}

// NewVote represent a new vote from the current user
// swagger:response vote
type NewVote struct {
	Vote int8 `json:"vote"`
}

// UserInformations represents the user information
// swagger:response userInfo
type UserInfo struct {
	Info     nerdz.InfoTO         `json:"info"`
	Contacts nerdz.ContactInfoTO  `json:"contacts"`
	Personal nerdz.PersonalInfoTO `json:"personal"`
}

// ProjectInfo represents the project information
// swagger:response projectInfo
type ProjectInfo struct {
	Info nerdz.InfoTO `json:"info"`
}
