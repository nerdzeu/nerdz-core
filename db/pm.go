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
	"errors"
	"time"

	"github.com/galeone/igor"
)

const (
	// MinPms represents the minimum pms number that can be required in a conversation
	MinPms uint64 = 1
	// MaxPms represents the maximum pms number that can be required in a conversation
	MaxPms uint64 = 20
)

// PmsOptions represent the configuration used to fetch a Pm list
type PmsOptions struct {
	N     uint8  // number of pms to return
	Older uint64 // if specified, tells to the function that is using this struct to return N pms OLDER (created before) than the pms with the specified "Older" ID
	Newer uint64 // if specified, tells to the function that is using this struct to return N pms NEWER (created after) than the comment with the spefified "Newer" ID
}

// pmsQueryBuilder returns the same pointer passed as first argument, with new specified options setted
func pmsQueryBuilder(query *igor.Database, options PmsOptions) *igor.Database {
	query = query.Limit(int(AtMostPms(uint64(options.N)))).Order("pmid DESC")
	if options.Older != 0 && options.Newer != 0 {
		query = query.Where("pmid BETWEEN ? AND ?", options.Newer, options.Older)
	} else if options.Older != 0 {
		query = query.Where("pmid < ?", options.Older)
	} else if options.Newer != 0 {
		query = query.Where("pmid > ?", options.Newer)
	}
	return query
}

// Conversation represents the details about a single private conversation between two users
type Conversation struct {
	From        uint64
	To          uint64
	LastMessage string
	Time        time.Time
	ToRead      bool
}

// NewPm initializes a Pm struct
func NewPm(pmid uint64) (*PM, error) {
	return NewPmWhere(&PM{Pmid: pmid})
}

// NewPmWhere returns the *Pm fetching the first one that matches the description
func NewPmWhere(description *PM) (pm *PM, e error) {
	pm = new(PM)
	if e = db().Model(PM{}).Where(description).Scan(pm); e != nil {
		return nil, e
	}
	if pm.Pmid == 0 {
		return nil, errors.New("Requested Pm does not exist")
	}
	return
}

// Implementing newMessage interface

// SetSender sets the source of the pm (the user ID)
func (pm *PM) SetSender(id uint64) {
	pm.From = id
}

// SetReference sets the destionation of the pm: user ID
func (pm *PM) SetReference(id uint64) {
	pm.To = id
}

// SetText set the text of the message
func (pm *PM) SetText(message string) {
	pm.Message = message
}

// SetLanguage set the language of the pm (useless)
func (pm *PM) SetLanguage(language string) error {
	pm.Lang = language
	return nil
}

// ClearDefaults set to the go's default values the fields with default sql values
func (pm *PM) ClearDefaults() {
	pm.Time = time.Time{}
}

// Implementing existingMessage interface

// ID returns the User Post ID
func (pm *PM) ID() uint64 {
	return pm.Pmid
}

// Language returns the message language
func (pm *PM) Language() string {
	return pm.Lang
}

// NumericSender returns the id of the sender user
func (pm *PM) NumericSender() uint64 {
	return pm.From
}

// Sender returns the sender *User
func (pm *PM) Sender() *User {
	user, _ := NewUser(pm.NumericSender())
	return user
}

// NumericReference returns the id of the recipient user
func (pm *PM) NumericReference() uint64 {
	return pm.To
}

// Reference returns the recipient *User
func (pm *PM) Reference() Reference {
	user, _ := NewUser(pm.NumericReference())
	return user
}

// Text returns the pm message
func (pm *PM) Text() string {
	return pm.Message
}

// IsEditable returns true if the pm is editable
func (pm *PM) IsEditable() bool {
	return false
}

// NumericOwners returns a slice of ids of the owner of the pms (the ones that can perform actions)
func (pm *PM) NumericOwners() []uint64 {
	return []uint64{pm.To, pm.From}
}

// Owners returns a slice of *User representing the users who own the pm
func (pm *PM) Owners() (ret []*User) {
	return Users(pm.NumericOwners())
}

// Revisions returns all the revisions of the message
func (pm *PM) Revisions() (modifications []string) {
	return
}

// RevisionsNumber returns the number of the revisions
func (pm *PM) RevisionsNumber() uint8 {
	return 0
}

// Votes returns the pm's votes value
func (pm *PM) VotesCount() int {
	return 0
}

// Voters returns a slice of *Vote representing the votes
func (pm *PM) Votes() (votes *[]Vote) {
	return
}
