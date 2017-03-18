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

package db

import (
	"net/mail"
	"net/url"
	"time"
)

// PersonalInfo is the struct that contains all the personal info of an user
type PersonalInfo struct {
	IsOnline  bool
	Nation    string
	Timezone  string
	Username  string
	Name      string
	Surname   string
	Gender    bool
	Birthday  time.Time
	Gravatar  *url.URL
	Interests []string
	Quotes    []string
	Biography string
}

// ContactInfo is the struct that contains all the contact info of an user
type ContactInfo struct {
	Website  *url.URL
	GitHub   *url.URL
	Skype    string
	Jabber   string
	Yahoo    *mail.Address
	Facebook *url.URL
	Twitter  *url.URL
	Steam    string
}

// BoardInfo is that struct that contains all the informations related to the user's board
type BoardInfo struct {
	Language   string
	IsClosed   bool
	Private    bool
	Whitelist  []*User
	UserScript *url.URL
}

// Info contains the informations common to every board
// Used in API output to give user/project basic informations
type Info struct {
	ID       uint64
	Owner    *Info
	Name     string
	Username string
	Website  *url.URL
	Image    *url.URL
	Closed   bool
	Type     boardType
}
