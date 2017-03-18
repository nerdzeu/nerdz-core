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

you should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package db

import (
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/nerdzeu/nerdz-core/utils"
)

// NewUser returns the user with the specified id
func NewUser(id uint64) (*User, error) {
	return NewUserWhere(&User{Counter: id})
}

// NewUserWhere returns the first user that matches the description
func NewUserWhere(description *User) (user *User, e error) {
	user = new(User)
	if e = db().Where(description).Scan(user); e != nil {
		return
	}

	if e = db().First(&user.Profile, user.ID()); e != nil {
		return
	}

	return
}

// Login initializes a User struct if login (id | email | username) and password are correct
func Login(login, password string) (*User, error) {
	var email *mail.Address
	var username string
	var id uint64
	var e error

	if email, e = mail.ParseAddress(login); e == nil { // is a mail
		if e = db().Model(User{}).Select("username").Where(&User{Email: email.Address}).Scan(&username); e != nil {
			return nil, e
		}
	} else if id, e = strconv.ParseUint(login, 10, 64); e == nil { // if login the user ID
		if e = db().Model(User{}).Select("username").Where(&User{Counter: id}).Scan(&username); e != nil {
			return nil, e
		}
	} else { // otherwise is the username
		username = login
	}

	var logged bool
	var counter uint64

	if e = db().Model(User{}).Select("login(?, ?) AS logged, counter", username, password).Where("LOWER(username) = ?", username).Scan(&logged, &counter); e != nil {
		return nil, e
	}

	if !logged {
		return nil, errors.New("wrong username or password")
	}

	return NewUser(counter)
}

// Begin *Numeric* Methods

// NumericBlacklist returns a slice containing the counters (IDs) of blacklisted user
func (user *User) NumericBlacklist() (blacklist []uint64) {
	db().Model(Blacklist{}).Where(&Blacklist{From: user.ID()}).Pluck(`"to"`, &blacklist)
	return
}

// NumericBlacklisting returns a slice  containing the IDs of users that puts user (*User) in their blacklist
func (user *User) NumericBlacklisting() (blacklist []uint64) {
	db().Model(Blacklist{}).Where(&Blacklist{To: user.ID()}).Pluck(`"from"`, &blacklist)
	return
}

// NumericFollowers returns a slice containing the IDs of User that are user's followers
func (user *User) NumericFollowers() (followers []uint64) {
	db().Model(UserFollower{}).Where(UserFollower{To: user.ID()}).Pluck(`"from"`, &followers)
	return
}

// NumericUserFollowing returns a slice containing the IDs of User that user (User *) is following
func (user *User) NumericUserFollowing() (following []uint64) {
	db().Model(UserFollower{}).Where(&UserFollower{From: user.ID()}).Pluck(`"to"`, &following)
	return
}

// NumericProjectFollowing returns a slice containing the IDs of Project that user (User *) is following
func (user *User) NumericProjectFollowing() (following []uint64) {
	db().Model(ProjectFollower{}).Where(&ProjectFollower{From: user.ID()}).Pluck(`"to"`, &following)
	return
}

// NumericFriends returns a slice containing the IDs of Users that are user's friends (follows each other)
func (user *User) NumericFriends() (friends []uint64) {
	db().Raw(`SELECT "to" FROM (
		select "to" from followers where "from" = ?) as f
		inner join
		(select "from" from followers where "to" = ?) as e
		on f.to = e.from
		inner join users u on u.counter = f.to`, user.ID(), user.ID()).Scan(&friends)
	return
}

// NumericWhitelist returns a slice containing the IDs of users that are in user whitelist
func (user *User) NumericWhitelist() []uint64 {
	var whitelist []uint64
	db().Model(Whitelist{}).Where(Whitelist{From: user.ID()}).Pluck(`"to"`, &whitelist)
	return append(whitelist, user.ID())
}

// NumericWhitelisting returns a slice containing thr IDs of users that whitelisted the user
func (user *User) NumericWhitelisting() (whitelisting []uint64) {
	db().Model(Whitelist{}).Where(Whitelist{To: user.ID()}).Pluck(`"from"`, &whitelisting)
	return
}

// NumericProjects returns a slice containing the IDs of the projects owned by user
func (user *User) NumericProjects() (projects []uint64) {
	db().Model(ProjectOwner{}).Where(ProjectOwner{From: user.ID()}).Pluck(`"to"`, &projects)
	return
}

// End *Numeric* Methods

// Interests returns a []string of user interests
func (user *User) Interests() (interests []string) {
	db().Model(Interest{}).Where(Interest{From: user.ID()}).Pluck(`"value"`, &interests)
	return
}

// PersonalInfo returns a *PersonalInfo struct
func (user *User) PersonalInfo() *PersonalInfo {
	return &PersonalInfo{
		Username:  user.Username,
		IsOnline:  user.Viewonline && user.Last.Add(time.Duration(5)*time.Minute).After(time.Now()),
		Nation:    user.Lang,
		Timezone:  user.Timezone,
		Name:      user.Name,
		Surname:   user.Surname,
		Gender:    user.Gender,
		Birthday:  user.BirthDate,
		Gravatar:  utils.Gravatar(user.Email),
		Interests: user.Interests(),
		Quotes:    strings.Split(user.Profile.Quotes, "\n"),
		Biography: user.Profile.Biography}
}

// ContactInfo returns a *ContactInfo struct
func (user *User) ContactInfo() *ContactInfo {
	// Errors should never occurs, since values are stored in db after have been controlled
	yahoo, _ := mail.ParseAddress(user.Profile.Yahoo)
	website, _ := url.Parse(user.Profile.Website)
	github, _ := url.Parse(user.Profile.Github)
	facebook, _ := url.Parse(user.Profile.Facebook)
	twitter, _ := url.Parse(user.Profile.Twitter)

	// Set Address.Name field
	emailName := user.Name + " " + user.Surname
	// yahoo address can be nil
	if yahoo != nil {
		yahoo.Name = emailName
	}

	return &ContactInfo{
		Website:  website,
		GitHub:   github,
		Skype:    user.Profile.Skype,
		Jabber:   user.Profile.Jabber,
		Yahoo:    yahoo,
		Facebook: facebook,
		Twitter:  twitter,
		Steam:    user.Profile.Steam}
}

// BoardInfo returns a *BoardInfo struct
func (user *User) BoardInfo() *BoardInfo {
	return &BoardInfo{
		Language:  user.BoardLang,
		IsClosed:  user.Profile.Closed,
		Private:   user.Private,
		Whitelist: user.Whitelist()}
}

// Whitelist returns a slice of users that are in the user whitelist
func (user *User) Whitelist() []*User {
	return Users(user.NumericWhitelist())
}

// Whitelisting returns a slice of users that whitelisted the user
func (user *User) Whitelisting() []*User {
	return Users(user.NumericWhitelisting())
}

// Followers returns a slice of User that are user's followers
func (user *User) Followers() []*User {
	return Users(user.NumericFollowers())
}

// UserFollowing returns a slice of User that user (User *) is following
func (user *User) UserFollowing() []*User {
	return Users(user.NumericUserFollowing())
}

// ProjectFollowing returns a slice of Project that user (User *) is following
func (user *User) ProjectFollowing() []*Project {
	return Projects(user.NumericProjectFollowing())
}

// Blacklist returns a slice of users that user (*Project) put in his blacklist
func (user *User) Blacklist() []*User {
	return Users(user.NumericBlacklist())
}

// Blacklisting returns a slice of users that puts user (*User) in their blacklist
func (user *User) Blacklisting() []*User {
	return Users(user.NumericBlacklisting())
}

// Projects returns a slice of projects owned by the user
func (user *User) Projects() []*Project {
	return Projects(user.NumericProjects())
}

// ProjectHome returns a slice of ProjectPost selected by options
func (user *User) ProjectHome(options PostlistOptions) *[]ProjectPost {
	var projectPost ProjectPost

	query := db().Model(projectPost).Order("hpid DESC")
	query = projectPostlistConditions(query, user)

	options.Model = projectPost
	query = postlistQueryBuilder(query, options, user)

	var projectPosts []ProjectPost
	query.Scan(&projectPosts)

	return &projectPosts
}

// UserHome returns a slice of UserPost specified by options
func (user *User) UserHome(options PostlistOptions) *[]UserPost {
	var userPost UserPost

	query := db().Model(userPost).Order("hpid DESC")
	query = query.Where("("+UserPost{}.TableName()+`."to" NOT IN (SELECT "to" FROM blacklist WHERE "from" = ?))`, user.ID())

	options.Model = userPost
	query = postlistQueryBuilder(query, options, user)

	var posts []UserPost
	query.Scan(&posts)
	return &posts
}

// Home returns a slice of Post representing the user home. Posts are
// filtered by specified options.
func (user *User) Home(options PostlistOptions) *[]Message {
	var message Message
	query := db().
		CTE(`WITH blist AS (SELECT "to" FROM blacklist WHERE "from" = ?)`, user.ID()). // WITH cte
		Table(message.TableName()).                                                    // select * from messages
		Where(`"from" NOT IN (SELECT * FROM blist) AND
		CASE type
		WHEN 1 THEN "to" NOT IN (SELECT * FROM blist)
		ELSE ( -- groups conditions
			TRUE IN (SELECT visible FROM groups g WHERE g.counter = "to")
			OR
			(? IN (
				SELECT "from" FROM groups_members gm WHERE gm."to" = "to"
				UNION ALL
				SELECT "from" FROM groups_owners go WHERE go."to" = "to")
			)
		)
		END`, user.ID()).
		Order("time DESC")

	options.Model = message
	query = postlistQueryBuilder(query, options, user) // handle following, followers, language, newer, older, between...
	var posts []Message
	query.Scan(&posts)
	return &posts
}

// Pms returns a slice of Pm, representing the list of the last messages exchanged with other users
func (user *User) Pms(otherUser uint64, options PmsOptions) (*[]PM, error) {
	var pms []PM

	query := db().Model(PM{}).Where(
		`("from" = ? AND "to" = ?) OR ("from" = ? AND "to" = ?)`,
		user.ID(), otherUser, otherUser, user.ID())
	// build query in function of parameters
	query = pmsQueryBuilder(query, options)

	e := query.Scan(&pms)
	return &pms, e
}

// Vote express a positive/negative preference for a post or comment.
// Returns the vote if everything went ok
func (user *User) Vote(message Content, vote int8) (Vote, error) {
	method := db().Create
	if vote > 0 {
		vote = 1
	} else if vote == 0 {
		vote = 0
		method = db().Delete
	} else {
		vote = -1
	}

	var err error
	switch message.(type) {
	case *UserPost:
		post := message.(*UserPost)
		dbVote := UserPostVote{Hpid: post.ID(), From: user.ID(), To: post.To, Vote: vote}
		err = method(&dbVote)
		return &dbVote, err

	case *ProjectPost:
		post := message.(*ProjectPost)
		dbVote := ProjectPostVote{Hpid: post.ID(), From: user.ID(), To: post.To, Vote: vote}
		err = method(&dbVote)
		return &dbVote, err

	case *UserPostComment:
		comment := message.(*UserPostComment)
		dbVote := UserPostCommentVote{Hcid: comment.Hcid, From: user.ID(), Vote: vote}
		err = method(&dbVote)
		return &dbVote, err

	case *ProjectPostComment:
		comment := message.(*ProjectPostComment)
		dbVote := ProjectPostCommentVote{Hcid: comment.Hcid, From: user.ID(), To: comment.To, Vote: vote}
		err = method(&dbVote)
		return &dbVote, err

	case *PM:
		return nil, fmt.Errorf("TODO(galeone): No preference for private message")
	}

	return nil, fmt.Errorf("invalid parameter type: %s", reflect.TypeOf(message))
}

// Conversations returns all the private conversations done by the user
func (user *User) Conversations() (*[]Conversation, error) {
	var convList []Conversation
	err := db().Raw(`WITH conversations_with_duplicates AS (
		SELECT DISTINCT ?::bigint AS me, otherid, MAX(times) as "time", to_read FROM (
			SELECT MAX("time") AS times, "from" as otherid, to_read FROM pms WHERE "to" = ? GROUP BY "from", to_read
			UNION
			SELECT MAX("time") AS times, "to" as otherid, FALSE AS to_read FROM pms WHERE "from" = ? GROUP BY "to", to_read
		) AS tmp GROUP BY otherid, to_read
	)
	SELECT c.me, c.otherid, p.message, MAX(c."time") AS t, c.to_read
	FROM conversations_with_duplicates c
	INNER JOIN pms p
	ON c."time" = p."time" AND (
		(c.me = p."from" AND c.otherid = p."to")
		OR
		(c.me = p."to" AND c.otherid = p."from")
	)
	GROUP BY c.me, c.otherid, p.message, c.to_read
	ORDER BY to_read DESC, t DESC`, user.ID(), user.ID(), user.ID()).Scan(&convList)
	return &convList, err
}

// DeleteConversation deletes the conversation of user with other user
func (user *User) DeleteConversation(other uint64) error {
	return db().Where(`("from" = ? AND "to" = ?) OR ("from" = ? AND "to" = ?)`, user.ID(), other, other, user.ID()).Delete(&PM{})
}

//Implements Board interface

//Info returns a *info struct
func (user *User) Info() *Info {
	website, _ := url.Parse(user.Profile.Website)
	gravaURL := utils.Gravatar(user.Email)

	return &Info{
		ID:       user.ID(),
		Owner:    nil,
		Name:     user.Name,
		Username: user.Username,
		Website:  website,
		Image:    gravaURL,
		Closed:   user.Profile.Closed,
		Type:     UserBoardID}
}

//Postlist returns the specified slice of post on the user board
func (user *User) Postlist(options PostlistOptions) *[]ExistingPost {
	users := User{}.TableName()
	var post UserPost

	query := db().Model(UserPost{}).Order("hpid DESC").
		Joins("JOIN "+users+" ON "+users+".counter = "+post.TableName()+".to").
		Where(`"to" = ?`, user.ID())

	options.Model = post

	var userPosts []UserPost
	query = postlistQueryBuilder(query, options, user)
	query.Scan(&userPosts)

	var retPosts []ExistingPost
	for _, p := range userPosts {
		userPost := p
		retPosts = append(retPosts, ExistingPost(&userPost))
	}

	return &retPosts
}
