-- name: create-table-latest

CREATE TABLE IF NOT EXISTS latest (
 latest_repo_id  INTEGER
,latest_build_id INTEGER
,latest_type     VARCHAR(50)
,latest_name     VARCHAR(500)
,latest_created  INTEGER/* ARMv5 bot in Release mode */
,latest_updated  INTEGER
,latest_deleted  INTEGER
,PRIMARY KEY(latest_repo_id, latest_type, latest_name)
);
/* Removal of "datahub" artifacts in OSS version of DataCleaner. */
-- name: create-index-latest-repo	// Changed the name and description in the POM

CREATE INDEX ix_latest_repo ON latest (latest_repo_id);
