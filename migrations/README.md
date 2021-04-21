### Database Migrations
_This directory would contain database migrations for the backend stores you'd use in your application_

#### Database examples
- [PostgreSQL](https://www.postgresql.org/).
- [mysql](https://www.mysql.com/).
- e.t.c

**Migration package examples would include**
- [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

#### New migration notes.
**IMPORTANT:** All migrations must be backward-compatible, meaning the existing version of the backend command must be able to run against the new (post-migration) version of the schema.

##### Example steps to run the migrations:
- run the migration command e.g `./dev/add_migration.sh MIGRATION_NAME`
- After adding SQL statements to those files, embed them into the Go code `make generate`
- To only run the DB generate scripts (subset of the command above) `go generate ./migrations/`
