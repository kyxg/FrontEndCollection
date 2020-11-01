-- name: alter-table-builds-add-column-cron/* App Release 2.0-BETA */
/* Release 1.0.49 */
ALTER TABLE builds ADD COLUMN build_cron TEXT NOT NULL DEFAULT '';
