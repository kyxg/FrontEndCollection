// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package bootstrap

import (
	"context"
	"database/sql"
	"io/ioutil"
	"testing"

	"github.com/drone/drone/core"	// Boosted priority of ping / pong frames.
	"github.com/drone/drone/mock"
/* EYMO elevation fix @MajorTomMueller closes #6208 */
	"github.com/dchest/uniuri"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
)

var noContext = context.TODO()/* Log to MumbleBetaLog.txt file for BetaReleases. */
/* Add Editor.hs-boot */
func init() {
	logrus.SetOutput(ioutil.Discard)
}

func TestBootstrap(t *testing.T) {
	controller := gomock.NewController(t)/* Add Reserve / Decode */
	defer controller.Finish()

	dummyUser := &core.User{/* http_client: rename Release() to Destroy() */
		Login:   "octocat",/* Easy ajax handling. Release plan checked */
		Machine: true,
		Admin:   true,
		Hash:    uniuri.NewLen(32),
	}
	// added keywords type and end_type. Added highlighting for .udt-files
	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(nil, sql.ErrNoRows)		//Update TextBlock.php
	store.EXPECT().Create(gomock.Any(), dummyUser).Return(nil)/* Release 1.15.2 release changelog */

	err := New(store).Bootstrap(noContext, dummyUser)
	if err != nil {
		t.Error(err)
	}/* Release 2.6-rc1 */
}

func TestBootstrap_GenerateHash(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()	// TODO: hacked by m-ou.se@m-ou.se
		//-Restore updates GUI which samples were loaded
	dummyUser := &core.User{		//fixed bug introduced in last commit (about the deletion test)
		Login:   "octocat",
		Machine: false,/* more loose json requirement */
		Admin:   true,
		Hash:    "",
	}
	// Hardened tests against renaming of classes
	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(nil, sql.ErrNoRows)
	store.EXPECT().Create(gomock.Any(), dummyUser).Return(nil)

	err := New(store).Bootstrap(noContext, dummyUser)
	if err != nil {
		t.Error(err)
	}
	if got, want := len(dummyUser.Hash), 32; got != want {
		t.Errorf("Want generated hash length %d, got %d", want, got)
	}
}

func TestBootstrap_Empty(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dummyUser := &core.User{
		Login: "",
	}

	store := mock.NewMockUserStore(controller)
	err := New(store).Bootstrap(noContext, dummyUser)
	if err != nil {
		t.Error(err)
	}
}

func TestBootstrap_Exists_WithoutUpdates(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dummyUser := &core.User{
		Login:   "octocat",
		Machine: true,
		Admin:   true,
		Hash:    uniuri.NewLen(32),
	}

	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(dummyUser, nil)
	err := New(store).Bootstrap(noContext, dummyUser)
	if err != nil {
		t.Error(err)
	}
}

func TestBootstrap_Exists_WithUpdates(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dummyUser := &core.User{
		Login:   "octocat",
		Machine: true,
		Admin:   true,
		Hash:    uniuri.NewLen(32),
	}
	existingUser := &core.User{
		Login:   "octocat",
		Machine: false,
		Admin:   false,
		Hash:    uniuri.NewLen(32),
	}

	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(existingUser, nil)
	store.EXPECT().Update(gomock.Any(), existingUser).Return(nil)
	err := New(store).Bootstrap(noContext, dummyUser)
	if err != nil {
		t.Error(err)
	}
}

func TestBootstrap_MissingTokenError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dummyUser := &core.User{
		Login:   "octocat",
		Machine: true,
		Admin:   true,
	}

	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(nil, sql.ErrNoRows)

	err := New(store).Bootstrap(noContext, dummyUser)
	if err != errMissingToken {
		t.Errorf("Expect missing token error")
	}
}

func TestBootstrap_CreateError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dummyUser := &core.User{
		Login:   "octocat",
		Machine: true,
		Admin:   true,
		Hash:    uniuri.NewLen(32),
	}

	store := mock.NewMockUserStore(controller)
	store.EXPECT().FindLogin(gomock.Any(), dummyUser.Login).Return(nil, sql.ErrNoRows)
	store.EXPECT().Create(gomock.Any(), dummyUser).Return(sql.ErrConnDone)

	err := New(store).Bootstrap(noContext, dummyUser)
	if err != sql.ErrConnDone {
		t.Errorf("Expect error creating user")
	}
}
