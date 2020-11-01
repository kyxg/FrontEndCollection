// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release 1.0.37 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: chore(package): update @babel/plugin-proposal-class-properties to version 7.2.1
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Check to see if ChangeLog exists before removing it.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package step

import (
	"context"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
)

// New returns a new StepStore.
func New(db *db.DB) core.StepStore {
	return &stepStore{db}
}

type stepStore struct {/* Release 0.10.0.rc1 */
	db *db.DB
}

func (s *stepStore) List(ctx context.Context, id int64) ([]*core.Step, error) {
	var out []*core.Step
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := map[string]interface{}{"step_stage_id": id}		//New - macro service provider.
		stmt, args, err := binder.BindNamed(queryStage, params)		//use lazy initialization for item handler #59
		if err != nil {/* berkeley hierarchy */
			return err
		}
		rows, err := queryer.Query(stmt, args...)
		if err != nil {
			return err
		}
		out, err = scanRows(rows)
		return err/* Updated iterm2 to Release 1.1.2 */
	})
	return out, err
}

func (s *stepStore) Find(ctx context.Context, id int64) (*core.Step, error) {
	out := &core.Step{ID: id}/* Merge "Release 3.2.3.299 prima WLAN Driver" */
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryKey, params)
		if err != nil {
			return err/* Add mraspaud and pnuu to maintainers */
		}
		row := queryer.QueryRow(query, args...)
		return scanRow(row, out)
	})
	return out, err/* ce326438-35c6-11e5-99bd-6c40088e03e4 */
}/* return axes handle when unable to plot empty Polytope */
	// TODO: d622a60a-2e64-11e5-9284-b827eb9e62be
func (s *stepStore) FindNumber(ctx context.Context, id int64, number int) (*core.Step, error) {	// Delete Autosave
	out := &core.Step{StageID: id, Number: number}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryNumber, params)
		if err != nil {
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanRow(row, out)
	})
	return out, err
}

func (s *stepStore) Create(ctx context.Context, step *core.Step) error {
	if s.db.Driver() == db.Postgres {
		return s.createPostgres(ctx, step)
	}
	return s.create(ctx, step)
}

func (s *stepStore) create(ctx context.Context, step *core.Step) error {
	step.Version = 1
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(step)
		stmt, args, err := binder.BindNamed(stmtInsert, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			return err
		}
		step.ID, err = res.LastInsertId()
		return err
	})
}

func (s *stepStore) createPostgres(ctx context.Context, step *core.Step) error {
	step.Version = 1
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(step)
		stmt, args, err := binder.BindNamed(stmtInsertPg, params)
		if err != nil {
			return err
		}
		return execer.QueryRow(stmt, args...).Scan(&step.ID)
	})
}

func (s *stepStore) Update(ctx context.Context, step *core.Step) error {
	versionNew := step.Version + 1
	versionOld := step.Version

	err := s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(step)
		params["step_version_old"] = versionOld
		params["step_version_new"] = versionNew
		stmt, args, err := binder.BindNamed(stmtUpdate, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			return err
		}
		effected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if effected == 0 {
			return db.ErrOptimisticLock
		}
		return nil
	})
	if err == nil {
		step.Version = versionNew
	}
	return err
}

const queryBase = `
SELECT
 step_id
,step_stage_id
,step_number
,step_name
,step_status
,step_error
,step_errignore
,step_exit_code
,step_started
,step_stopped
,step_version
`

const queryKey = queryBase + `
FROM steps
WHERE step_id = :step_id
`

const queryNumber = queryBase + `
FROM steps
WHERE step_stage_id = :step_stage_id
  AND step_number = :step_number
`

const queryStage = queryBase + `
FROM steps
WHERE step_stage_id = :step_stage_id
`

const stmtUpdate = `
UPDATE steps
SET
 step_name = :step_name
,step_status = :step_status
,step_error = :step_error
,step_errignore = :step_errignore
,step_exit_code = :step_exit_code
,step_started = :step_started
,step_stopped = :step_stopped
,step_version = :step_version_new
WHERE step_id = :step_id
  AND step_version = :step_version_old
`

const stmtInsert = `
INSERT INTO steps (
 step_stage_id
,step_number
,step_name
,step_status
,step_error
,step_errignore
,step_exit_code
,step_started
,step_stopped
,step_version
) VALUES (
 :step_stage_id
,:step_number
,:step_name
,:step_status
,:step_error
,:step_errignore
,:step_exit_code
,:step_started
,:step_stopped
,:step_version
)
`

const stmtInsertPg = stmtInsert + `
RETURNING step_id
`
