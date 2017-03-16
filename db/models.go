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
	"database/sql"
	"github.com/galeone/igor"
	"time"
)

// Enrich models structure with unexported types

// boardType represents a board type
type boardType string

const (
	// UserBoardID constant (of type boardType) makes possible to distinguish a User
	// board from a Project board
	UserBoardID boardType = "user"
	// ProjectBoardID constant (of type boardType) makes possible to distinguish a PROJECT
	// board from a User board
	ProjectBoardID boardType = "project"
)

// Models

// UserPostLock is the model for the relation posts_no_notify
type UserPostLock struct {
	User    uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostLock) TableName() string {
	return "posts_no_notify"
}

// UserPostUserLock is the model for the relation comments_no_notify
type UserPostUserLock struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostUserLock) TableName() string {
	return "comments_no_notify"
}

// UserPostCommentsNotify is the model for the relation comments_notify
type UserPostCommentsNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostCommentsNotify) TableName() string {
	return "comments_notify"
}

// Ban is the model for the relation ban
type Ban struct {
	User       uint64
	Motivation string
	Time       time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter    uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (Ban) TableName() string {
	return "ban"
}

// Blacklist is the model for the relation blacklist
type Blacklist struct {
	From       uint64
	To         uint64
	Motivation string
	Time       time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter    uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (Blacklist) TableName() string {
	return "blacklist"
}

// Whitelist is the model for the relation whitelist
type Whitelist struct {
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (Whitelist) TableName() string {
	return "whitelist"
}

// UserFollower is the model for the relation followers
type UserFollower struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	ToNotify bool
	Counter  uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserFollower) TableName() string {
	return "followers"
}

// ProjectNotify is the model for the relation groups_notify
type ProjectNotify struct {
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Hpid    uint64
	Counter uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectNotify) TableName() string {
	return "groups_notify"
}

// ProjectPostLock is the model for the relation groups_posts_no_notify
type ProjectPostLock struct {
	User    uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectPostLock) TableName() string {
	return "groups_posts_no_notify"
}

// ProjectPostUserLock is the model for the relation groups_comments_no_notify
type ProjectPostUserLock struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectPostUserLock) TableName() string {
	return "groups_comments_no_notify"
}

// ProjectPostCommentsNotify is the model for the relation groups_comments_notify
type ProjectPostCommentsNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectPostCommentsNotify) TableName() string {
	return "groups_comments_notify"
}

// User is the model for the relation users
type User struct {
	Counter          uint64    `igor:"primary_key"`
	Last             time.Time `sql:"default:(now() at time zone 'utc')"`
	NotifyStory      igor.JSON `sql:"default:'{}'::jsonb"`
	Private          bool
	Lang             string
	Username         string
	Password         string
	RemoteAddr       string
	HTTPUserAgent    string `igor:"column:http_user_agent"`
	Email            string
	Name             string
	Surname          string
	Gender           bool
	BirthDate        time.Time `sql:"default:(now() at time zone 'utc')"`
	BoardLang        string
	Timezone         string
	Viewonline       bool
	RegistrationTime time.Time `sql:"default:(now() at time zone 'utc')"`
	// Relation. Manually fill the field when required
	Profile Profile `sql:"-"`
}

// TableName returns the table name associated with the structure
func (User) TableName() string {
	return "users"
}

// Profile is the model for the relation profiles
type Profile struct {
	Counter        uint64 `igor:"primary_key"`
	Website        string
	Quotes         string
	Biography      string
	Github         string
	Skype          string
	Jabber         string
	Yahoo          string
	Userscript     string
	Template       uint8
	MobileTemplate uint8
	Dateformat     string
	Facebook       string
	Twitter        string
	Steam          string
	Push           bool
	Pushregtime    time.Time `sql:"default:(now() at time zone 'utc')"`
	Closed         bool
}

// TableName returns the table name associated with the structure
func (Profile) TableName() string {
	return "profiles"
}

// Interest is the model for the relation interests
type Interest struct {
	ID    uint64 `igor:"primary_key"`
	From  uint64
	Value string
	Time  time.Time `sql:"default:(now() at time zone 'utc')"`
}

// TableName returns the table name associated with the structure
func (Interest) TableName() string {
	return "interests"
}

// Post is the type of a generic post
type Post struct {
	Hpid    uint64 `igor:"primary_key"`
	From    uint64
	To      uint64
	Pid     uint64 `sql:"default:0"`
	Message string
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Lang    string
	News    bool
	Closed  bool
}

// UserPost converts the Post to a UserPost
func (p *Post) UserPost() *UserPost {
	return &UserPost{
		Post: *p,
	}
}

// ProjectPost converts the Post to ProjectPost
func (p *Post) ProjectPost() *ProjectPost {
	return &ProjectPost{
		Post: *p,
	}
}

// UserPost is the model for the relation posts
type UserPost struct {
	Post
}

// TableName returns the table name associated with the structure
func (UserPost) TableName() string {
	return "posts"
}

// UserPostRevision is the model for the relation posts_revisions
type UserPostRevision struct {
	Hpid    uint64
	Message string
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	RevNo   uint16
	Counter uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostRevision) TableName() string {
	return "posts_revisions"
}

// UserPostVote is the model for the relation votes
type UserPostVote struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Vote    int8
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostVote) TableName() string {
	return "thumbs"
}

// UserPostLurk is the model for the relation lurkers
type UserPostLurk struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostLurk) TableName() string {
	return "lurkers"
}

// UserPostComment is the model for the relation comments
type UserPostComment struct {
	Hcid     uint64 `igor:"primary_key"`
	Hpid     uint64
	From     uint64
	To       uint64
	Message  string
	Lang     string
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	Editable bool      `sql:"default:true"`
}

// TableName returns the table name associated with the structure
func (UserPostComment) TableName() string {
	return "comments"
}

// UserPostCommentRevision is the model for the relation comments_revisions
type UserPostCommentRevision struct {
	Hcid    uint64
	Message string
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	RevNo   int8
	Counter uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostCommentRevision) TableName() string {
	return "comments_revisions"
}

// UserPostBookmark is the model for the relation bookmarks
type UserPostBookmark struct {
	Hpid    uint64
	From    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostBookmark) TableName() string {
	return "bookmarks"
}

// Pm is the model for the relation pms
type Pm struct {
	Pmid    uint64 `igor:"primary_key"`
	From    uint64
	To      uint64
	Message string
	Lang    string
	ToRead  bool
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
}

// TableName returns the table name associated with the structure
func (Pm) TableName() string {
	return "pms"
}

// Project is the model for the relation groups
type Project struct {
	Counter      uint64 `igor:"primary_key"`
	Description  string
	Name         string
	Private      bool
	Photo        sql.NullString
	Website      sql.NullString
	Goal         string
	Visible      bool
	Open         bool
	CreationTime time.Time `sql:"default:(now() at time zone 'utc')"`
}

// TableName returns the table name associated with the structure
func (Project) TableName() string {
	return "groups"
}

// ProjectMember is the model for the relation groups_members
type ProjectMember struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	ToNotify bool
	Counter  uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectMember) TableName() string {
	return "groups_members"
}

// ProjectOwner is the model for the relation groups_owners
type ProjectOwner struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	ToNotify bool
	Counter  uint64 `igor:"primary_key"`
}

// GetTO returns its Transfer Object
// TableName returns the table name associated with the structure
func (ProjectOwner) TableName() string {
	return "groups_owners"
}

// ProjectPost is the model for the relation groups_posts
type ProjectPost struct {
	Post
}

// TableName returns the table name associated with the structure
func (ProjectPost) TableName() string {
	return "groups_posts"
}

// ProjectPostRevision is the model for the relation groups_posts_revisions
type ProjectPostRevision struct {
	Hpid    uint64
	Message string
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	RevNo   uint16
	Counter uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectPostRevision) TableName() string {
	return "groups_posts_revisions"
}

// ProjectPostVote is the model for the relation groups_thumbs
type ProjectPostVote struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Vote    int8
	Counter uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectPostVote) TableName() string {
	return "groups_thumbs"
}

// ProjectPostLurk is the model for the relation groups_lurkers
type ProjectPostLurk struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectPostLurk) TableName() string {
	return "groups_lurkers"
}

// ProjectPostComment is the model for the relation groups_comments
type ProjectPostComment struct {
	Hcid     uint64 `igor:"primary_key"`
	Hpid     uint64
	From     uint64
	To       uint64
	Message  string
	Lang     string
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	Editable bool      `sql:"default:true"`
}

// TableName returns the table name associated with the structure
func (ProjectPostComment) TableName() string {
	return "groups_comments"
}

// ProjectPostCommentRevision is the model for the relation groups_comments_revisions
type ProjectPostCommentRevision struct {
	Hcid    uint64
	Message string
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	RevNo   uint16
	Counter uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectPostCommentRevision) TableName() string {
	return "groups_comments_revisions"
}

// ProjectPostBookmark is the model for the relation groups_bookmarks
type ProjectPostBookmark struct {
	Hpid    uint64
	From    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectPostBookmark) TableName() string {
	return "groups_bookmarks"
}

// ProjectFollower is the model for the relation groups_followers
type ProjectFollower struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	ToNotify bool
	Counter  uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectFollower) TableName() string {
	return "groups_followers"
}

// UserPostCommentVote is the model for the relation groups_comment_thumbs
type UserPostCommentVote struct {
	Hcid    uint64
	From    uint64
	Vote    int8
	Counter uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostCommentVote) TableName() string {
	return "comment_thumbs"
}

// ProjectPostCommentVote is the model for the relation groups_comment_thumbs
type ProjectPostCommentVote struct {
	Hcid    uint64
	From    uint64
	To      uint64
	Vote    int8
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectPostCommentVote) TableName() string {
	return "groups_comment_thumbs"
}

// DeletedUser is the model for the relation deleted_users
type DeletedUser struct {
	Counter    uint64 `igor:"primary_key"`
	Username   string
	Time       time.Time `sql:"default:(now() at time zone 'utc')"`
	Motivation string
}

// TableName returns the table name associated with the structure
func (DeletedUser) TableName() string {
	return "deleted_users"
}

// SpecialUser is the model for the relation special_users
type SpecialUser struct {
	Role    string `igor:"primary_key"`
	Counter uint64
}

// TableName returns the table name associated with the structure
func (SpecialUser) TableName() string {
	return "special_users"
}

// SpecialProject is the model for the relation special_groups
type SpecialProject struct {
	Role    string `igor:"primary_key"`
	Counter uint64
}

// TableName returns the table name associated with the structure
func (SpecialProject) TableName() string {
	return "special_groups"
}

// PostClassification is the model for the relation posts_classifications
type PostClassification struct {
	ID    uint64 `igor:"primary_key"`
	UHpid uint64
	GHpid uint64
	Tag   string
}

// TableName returns the table name associated with the structure
func (PostClassification) TableName() string {
	return "posts_classifications"
}

// Mention is the model for the relation mentions
type Mention struct {
	ID       uint64 `igor:"primary_key"`
	UHpid    uint64
	GHpid    uint64
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	ToNotify bool
}

// TableName returns the table name associated with the structure
func (Mention) TableName() string {
	return "mentions"
}

// Message is the model for the view message
type Message struct {
	Post
	Type uint8
}

// TableName returns the table name associated with the structure
func (Message) TableName() string {
	return "messages"
}

// OAuth2Client implements the osin.Client interface
type OAuth2Client struct {
	// Surrogated key
	ID uint64 `igor:"primary_key"`
	// Real Primary Key. Application (client) name
	Name string `sql:"UNIQUE"`
	// Secret is the unique secret associated with a client
	Secret string `sql:"UNIQUE"`
	// RedirectURI is the valid redirection URI associated with a client
	RedirectURI string
	// UserID references User that created this client
	UserID uint64
}

// TableName returns the table name associated with the structure
func (OAuth2Client) TableName() string {
	return "oauth2_clients"
}

// OAuth2AuthorizeData is the model for the relation oauth2_authorize
// that represents the authorization granted to to the client
type OAuth2AuthorizeData struct {
	// Surrogated key
	ID uint64 `igor:"primary_key"`
	// ClientID references the client that created this token
	ClientID uint64
	// Code is the Authorization code
	Code string
	// CreatedAt is the instant of creation of the OAuth2AuthorizeToken
	CreatedAt time.Time `sql:"default:(now() at time zone 'utc')"`
	// ExpiresIn is the seconds from CreatedAt before this token expires
	ExpiresIn uint64
	// State data from request
	//State string, [!] we dont't store state variables
	// Scope is the requested scope
	Scope string
	// RedirectUri is the RedirectUri associated with the token
	RedirectURI string
	// UserID is references the User that created the authorization request and thus the AuthorizeData
	UserID uint64
}

// TableName returns the table name associated with the structure
func (OAuth2AuthorizeData) TableName() string {
	return "oauth2_authorize"
}

// OAuth2AccessData is the OAuth2 access data
type OAuth2AccessData struct {
	ID uint64 `igor:"primary_key"`
	// ClientID references the client that created this token
	ClientID uint64
	// CreatedAt is the instant of creation of the OAuth2AccessToken
	CreatedAt time.Time `sql:"default:(now() at time zone 'utc')"`
	// ExpiresIn is the seconds from CreatedAt before this token expires
	ExpiresIn uint64
	// RedirectUri is the RedirectUri associated with the token
	RedirectURI string
	// AuthorizeDataID references the AuthorizationData that authorizated this token. Can be null
	AuthorizeDataID sql.NullInt64 `igor:"column:oauth2_authorize_id"` // Annotation required, since the column name does not follow igor conventions
	// AccessDataID references the Access Data, for refresh token. Can be null
	AccessDataID sql.NullInt64 `igor:"column:oauth2_access_id"` // Annotation required, since the column name does not follow igor conventions
	// RefreshTokenID is the value by which this token can be renewed. Can be null
	RefreshTokenID sql.NullInt64
	// AccessToken is the main value of this tructure, represents the access token
	AccessToken string
	// Scope is the requested scope
	Scope string
	// UserID is references the User that created The access request and thus the AccessData
	UserID uint64
}

// TableName returns the table name associated with the structure
func (OAuth2AccessData) TableName() string {
	return "oauth2_access"
}

// OAuth2RefreshToken is the model for the relation oauth2_refresh
type OAuth2RefreshToken struct {
	ID    uint64 `igor:"primary_key"`
	Token string `sql:"UNIQUE"`
}

// TableName returns the table name associated with the structure
func (OAuth2RefreshToken) TableName() string {
	return "oauth2_refresh"
}
