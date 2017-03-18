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
	"reflect"

	"html"

	"github.com/galeone/igor"
	"github.com/nerdzeu/nerdz-core/utils"
)

// User actions

// Delete an existing message
func (user *User) Delete(message Content) error {
	if user.CanDelete(message) {
		return db().Delete(message)
	}
	return errors.New("you can't delete this message")
}

// Edit an existing message
func (user *User) Edit(message Content) error {
	if user.CanEdit(message) {
		rollBackText := message.Text() //unencoded

		if err := populateContent(message, user); err != nil {
			message.SetText(rollBackText)
			return err
		}

		if err := db().Updates(message); err != nil {
			message.SetText(rollBackText)
			return err
		}

		return nil
	}

	return errors.New("editing of this message is not allowed")
}

// Follow creates a new "follow" relationship between the current user
// and another NERDZ board. The board could represent a NERDZ's project
// or another NERDZ's user.
func (user *User) Follow(board Board) error {
	if board == nil {
		return errors.New("unable to follow an undefined board")
	}

	switch board.(type) {
	case *User:
		otherUser := board.(*User)
		return db().Create(&UserFollower{From: user.ID(), To: otherUser.ID()})

	case *Project:
		otherProj := board.(*Project)
		return db().Create(&ProjectFollower{From: user.ID(), To: otherProj.ID()})

	}

	return errors.New("invalid follower type " + reflect.TypeOf(board).String())
}

// Submit submits a Message
func (user *User) Submit(message Content) error {
	if err := populateContent(message, user); err != nil {
		return err
	}

	return db().Create(message.(igor.DBModel))
}

// WhitelistUser add other user to the user whitelist
func (user *User) WhitelistUser(other *User) error {
	if other == nil {
		return errors.New("Other user should be a vaid user")
	}

	return db().Create(&Whitelist{From: user.ID(), To: other.ID()})
}

// UnwhitelistUser removes other user to the user whitelist
func (user *User) UnwhitelistUser(other *User) error {
	if other == nil {
		return errors.New("Other user should be a vaid user")
	}

	return db().Where(&Whitelist{From: user.ID(), To: other.ID()}).Delete(Whitelist{})
}

// BlacklistUser add other user to the user blacklist
func (user *User) BlacklistUser(other *User, motivation string) error {
	if other == nil {
		return errors.New("Other user should be a vaid user")
	}
	return db().Create(&Blacklist{From: user.ID(), To: other.ID(), Motivation: motivation})
}

// UnblacklistUser removes other user to the user blacklist
func (user *User) UnblacklistUser(other *User) error {
	if other == nil {
		return errors.New("Other user should be a vaid user")
	}
	return db().Where(&Blacklist{From: user.ID(), To: other.ID()}).Delete(Blacklist{})
}

// Unfollow delete a "follow" relationship between the current user
// and another NERDZ board. The board could represent a NERDZ's project
// or another NERDZ's user.
func (user *User) Unfollow(board Board) error {
	if board == nil {
		return errors.New("unable to unfollow an undefined board")
	}

	switch board.(type) {
	case *User:
		otherUser := board.(*User)
		return db().Where(&UserFollower{From: user.ID(), To: otherUser.ID()}).Delete(UserFollower{})

	case *Project:
		otherProj := board.(*Project)
		return db().Where(&ProjectFollower{From: user.ID(), To: otherProj.ID()}).Delete(ProjectFollower{})

	}

	return errors.New("invalid follower type " + reflect.TypeOf(board).String())
}

// Bookmark bookmarks the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the
// DBMS
func (user *User) Bookmark(post ExistingPost) (Bookmark, error) {
	if post == nil {
		return nil, errors.New("unable to bookmark undefined post")
	}

	switch post.(type) {
	case *UserPost:
		userPost := post.(*UserPost)
		bookmark := UserPostBookmark{From: user.ID(), Hpid: userPost.ID()}
		err := db().Create(&bookmark)
		return &bookmark, err

	case *ProjectPost:
		projectPost := post.(*ProjectPost)
		bookmark := ProjectPostBookmark{From: user.ID(), Hpid: projectPost.ID()}
		err := db().Create(&bookmark)
		return &bookmark, err
	}

	return nil, errors.New("invalid post type " + reflect.TypeOf(post).String())
}

// Unbookmark the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the DBMS
func (user *User) Unbookmark(post ExistingPost) error {
	if post == nil {
		return errors.New("unable to unbookmark undefined post")
	}

	switch post.(type) {
	case *UserPost:
		userPost := post.(*UserPost)
		return db().Where(&UserPostBookmark{From: user.ID(), Hpid: userPost.ID()}).Delete(UserPostBookmark{})

	case *ProjectPost:
		projectPost := post.(*ProjectPost)
		return db().Where(&ProjectPostBookmark{From: user.ID(), Hpid: projectPost.ID()}).Delete(ProjectPostBookmark{})
	}

	return errors.New("invalid post type " + reflect.TypeOf(post).String())
}

// Lurk lurkes the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the
// DBMS
func (user *User) Lurk(post ExistingPost) (Lurk, error) {
	if post == nil {
		return nil, errors.New("unable to lurk undefined post")
	}

	switch post.(type) {
	case *UserPost:
		userPost := post.(*UserPost)
		lurk := UserPostLurk{From: user.ID(), Hpid: userPost.ID()}
		err := db().Create(&lurk)
		return &lurk, err

	case *ProjectPost:
		projectPost := post.(*ProjectPost)
		lurk := ProjectPostLurk{From: user.ID(), Hpid: projectPost.ID()}
		err := db().Create(&lurk)
		return &lurk, err
	}

	return nil, errors.New("invalid post type " + reflect.TypeOf(post).String())
}

// Unlurk the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the DBMS
func (user *User) Unlurk(post ExistingPost) error {
	if post == nil {
		return errors.New("unable to unlurk undefined post")
	}

	switch post.(type) {
	case *UserPost:
		userPost := post.(*UserPost)
		return db().Where(&UserPostLurk{From: user.ID(), Hpid: userPost.ID()}).Delete(UserPostLurk{})

	case *ProjectPost:
		projectPost := post.(*ProjectPost)
		return db().Where(&ProjectPostLurk{From: user.ID(), Hpid: projectPost.ID()}).Delete(ProjectPostLurk{})
	}

	return errors.New("invalid post type " + reflect.TypeOf(post).String())
}

// LockPost lockes the specified post. If users are present, indiidual notifications
// are disabled from the user presents in the users list.
func (user *User) LockPost(post ExistingPost, users ...*User) (*[]Lock, error) {
	if post == nil {
		return nil, errors.New("unable to lurk undefined post")
	}

	switch post.(type) {
	case *UserPost:
		userPost := post.(*UserPost)
		if len(users) == 0 {
			lock := UserPostLock{User: user.ID(), Hpid: userPost.ID()}
			err := db().Create(&lock)
			return &[]Lock{&lock}, err
		}
		var locks []Lock
		for _, other := range users {
			lock := UserPostUserLock{From: user.ID(), To: other.ID(), Hpid: userPost.ID()}
			if err := db().Create(&lock); err != nil {
				return nil, err
			}
			locks = append(locks, Lock(&lock))
		}
		return &locks, nil

	case *ProjectPost:
		projectPost := post.(*ProjectPost)
		if len(users) == 0 {
			projectPost := post.(*ProjectPost)
			lock := ProjectPostLock{User: user.ID(), Hpid: projectPost.ID()}
			err := db().Create(&lock)
			return &[]Lock{&lock}, err
		}
		var locks []Lock
		for _, other := range users {
			lock := ProjectPostUserLock{From: user.ID(), To: other.ID(), Hpid: projectPost.ID()}
			if err := db().Create(&lock); err != nil {
				return nil, err
			}
			locks = append(locks, Lock(&lock))
		}
		return &locks, nil
	}

	return nil, errors.New("invalid post type " + reflect.TypeOf(post).String())
}

// Unlock the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the DBMS
func (user *User) Unlock(post ExistingPost, users ...*User) error {
	if post == nil {
		return errors.New("unable to unlock undefined post")
	}

	switch post.(type) {
	case *UserPost:
		userPost := post.(*UserPost)
		if len(users) == 0 {
			return db().Where(&UserPostLock{User: user.ID(), Hpid: userPost.ID()}).Delete(UserPostLock{})
		}
		for _, other := range users {
			err := db().Where(&UserPostUserLock{From: user.ID(), To: other.ID(), Hpid: userPost.ID()}).Delete(UserPostUserLock{})
			if err != nil {
				return err
			}
		}
		return nil

	case *ProjectPost:
		projectPost := post.(*ProjectPost)
		if len(users) == 0 {
			return db().Where(&ProjectPostLock{User: user.ID(), Hpid: projectPost.ID()}).Delete(ProjectPostLock{})
		}
		for _, other := range users {
			err := db().Where(&ProjectPostUserLock{From: user.ID(), To: other.ID(), Hpid: projectPost.ID()}).Delete(ProjectPostUserLock{})
			if err != nil {
				return err
			}
		}
		return nil
	}

	return errors.New("invalid post type " + reflect.TypeOf(post).String())
}

// AddInterest adds the specified interest. An error is returned if the
// interests already exists or some DBMS contraint is violated
func (user *User) AddInterest(interest *Interest) error {
	interest.From = user.ID()
	if interest.Value == "" {
		return errors.New("invalid interest value: (empty)")
	}
	return db().Create(interest)
}

// DeleteInterest removes the specified interest (by its ID or its Value).
func (user *User) DeleteInterest(interest *Interest) error {
	var toDelete Interest
	if interest.ID <= 0 {
		if interest.Value == "" {
			return errors.New("invalid interest ID and empty interest")
		}
		toDelete.Value = interest.Value
	} else {
		toDelete.ID = interest.ID
	}

	if interest.From != user.ID() {
		return errors.New("you can't remove other user interests")
	}

	toDelete.From = interest.From

	return db().Where(&toDelete).Delete(Interest{})
}

// Friends returns the current user's friends
func (user *User) Friends() []*User {
	return Users(user.NumericFriends())
}

// Implements Reference interface

// ID returns the user ID
func (user *User) ID() uint64 {
	return user.Counter
}

// Language returns the user language
func (user *User) Language() string {
	return user.Lang
}

// Can* methods

// CanEdit returns true if user can edit the Message
func (user *User) CanEdit(message Content) bool {
	return message.ID() > 0 && message.IsEditable() && utils.InSlice(user.ID(), message.NumericOwners())
}

// CanDelete returns true if user can delete the Message
func (user *User) CanDelete(message Content) bool {
	return message.ID() > 0 && utils.InSlice(user.ID(), message.NumericOwners())
}

// CanBookmark returns true if user haven't bookamrked to existingPost yet
func (user *User) CanBookmark(message ExistingPost) bool {
	return message.ID() > 0 && !utils.InSlice(user.ID(), message.NumericBookmarkers())
}

// CanLurk returns true if the user haven't lurked the existingPost yet
func (user *User) CanLurk(message ExistingPost) bool {
	return message.ID() > 0 && !utils.InSlice(user.ID(), message.NumericLurkers())
}

// CanComment returns true if the user can comment to the existingPost
func (user *User) CanComment(message ExistingPost) bool {
	return !utils.InSlice(user.ID(), message.Sender().NumericBlacklist()) && message.ID() > 0 && !message.IsClosed()
}

// CanSee returns true if the user can see the Board content
func (user *User) CanSee(board Board) bool {
	switch board.(type) {
	case *User:
		return !utils.InSlice(user.ID(), board.(*User).NumericBlacklist())

	case *Project:
		project := board.(*Project)
		if project.Visible {
			return true
		}

		return user.ID() == project.NumericOwner() || utils.InSlice(user.ID(), project.NumericMembers())
	}
	return false
}

// populate functions

func populateContent(message Content, user *User) error {
	message.ClearDefaults()

	if post, ok := message.(*UserPost); ok && post.To == 0 {
		post.To = user.ID()
	}

	lang, err := sanitiseLanguage(message.Language(), message.Sender().Language())
	if err != nil {
		return err
	}

	message.SetLanguage(lang)
	message.SetSender(user.ID())
	message.SetText(html.EscapeString(message.Text()))

	return nil
}
