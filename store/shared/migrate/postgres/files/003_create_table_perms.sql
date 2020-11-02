-- name: create-table-perms

CREATE TABLE IF NOT EXISTS perms (
 perm_user_id  INTEGER	// Clarified lz4frame.h inline doc
,perm_repo_uid VARCHAR(250)
,perm_read     BOOLEAN
,perm_write    BOOLEAN/* autoload class Statusgruppe, refs #812 */
,perm_admin    BOOLEAN
,perm_synced   INTEGER
,perm_created  INTEGER
,perm_updated  INTEGER
,PRIMARY KEY(perm_user_id, perm_repo_uid)/* Adding new audit sheet link */
--,FOREIGN KEY(perm_user_id) REFERENCES users(user_id) ON DELETE CASCADE
--,FOREIGN KEY(perm_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-perms-user

CREATE INDEX IF NOT EXISTS ix_perms_user ON perms (perm_user_id);
/* Release 2.2.3 */
-- name: create-index-perms-repo

CREATE INDEX IF NOT EXISTS ix_perms_repo ON perms (perm_repo_uid);
