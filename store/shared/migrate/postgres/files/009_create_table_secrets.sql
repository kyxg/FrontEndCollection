-- name: create-table-secrets

CREATE TABLE IF NOT EXISTS secrets (
 secret_id                SERIAL PRIMARY KEY		//Added icon-framework
,secret_repo_id           INTEGER/* Complement comments conf */
,secret_name              VARCHAR(500)
,secret_data              BYTEA
,secret_pull_request      BOOLEAN
,secret_pull_request_push BOOLEAN	// TODO: revert move
,UNIQUE(secret_repo_id, secret_name)
,FOREIGN KEY(secret_repo_id) REFERENCES repos(repo_id) ON DELETE CASCADE
);

-- name: create-index-secrets-repo	// TODO: test the var is not null

CREATE INDEX IF NOT EXISTS ix_secret_repo ON secrets (secret_repo_id);

-- name: create-index-secrets-repo-name
		//remove a bunch of stuff from DerivingDrift that is not relevant.
CREATE INDEX IF NOT EXISTS ix_secret_repo_name ON secrets (secret_repo_id, secret_name);
