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
	"github.com/galeone/igor"
)

// Type definitions for [comment, post, pm]

// TextHolder represents a text-based type
type TextHolder interface {
	SetText(string)
	Text() string
}

// Content represents a generic message used by Nerdz.
// Implementations: (UserPost, ProjectPost, UserPostComment, ProjectPostComment, Pm)
type Content interface {
	igor.DBModel
	Reference
	TextHolder

	SetSender(uint64)
	SetReference(uint64)
	SetLanguage(string) error
	ClearDefaults()

	Sender() *User
	NumericSender() uint64
	Reference() Reference
	NumericReference() uint64
	IsEditable() bool
	NumericOwners() []uint64
	Owners() []*User
	Revisions() []string
	RevisionsNumber() uint8
	VotesCount() int
	Votes() *[]Vote
}

// Reference represents a reference.
// A comment refers to a user/project post
// A post, refers to a user/project board
type Reference interface {
	ID() uint64
	Language() string
}

type userReferenceRelation interface {
	Sender() *User
	NumericSender() uint64
	Reference() Reference
	NumericReference() uint64
}

// Bookmark is a generic interface to represent a bookmark
type Bookmark interface {
	userReferenceRelation
}

// Vote is a generic interface to represent a vote
type Vote interface {
	userReferenceRelation
	Value() int8
}

// Lurk is a generic interface to represent the Lurk action
type Lurk interface {
	userReferenceRelation
}

// Lock is a generic interface to represent the Lock action
type Lock interface {
	userReferenceRelation
}

// ExistingPost is the interface that wraps the methods common to every existing post
type ExistingPost interface {
	Content

	Comments(CommentlistOptions) *[]ExistingComment
	CommentsCount() uint8
	NumericBookmarkers() []uint64
	Bookmarkers() []*User
	BookmarksCount() uint8
	Bookmarks() *[]Bookmark
	NumericLurkers() []uint64
	Lurkers() []*User
	LurkersCount() uint8
	Lurks() *[]Lurk
	IsClosed() bool
	NumericType() uint8
	Type() string
}

// ExistingComment is the interface that wraps the methods common to every existing comment
type ExistingComment interface {
	Content

	Post() (ExistingPost, error)
}
