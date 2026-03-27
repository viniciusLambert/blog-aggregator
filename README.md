# Gator - RSS Feed Aggregator

A CLI tool for aggregating and browsing RSS feeds, built with Go and PostgreSQL.

## Prerequisites

- [Go](https://go.dev/doc/install) 1.22+
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

```bash
go install github.com/viniciusLambert/blog-aggregator@latest
```

This compiles and installs the `blog-aggregator` binary to your `$GOPATH/bin`. Make sure that directory is in your `PATH`.

## Configuration

Gator reads its config from `~/.gatorconfig.json`. Create the file with the following structure:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```

Replace `username`, `password`, and `gator` with your PostgreSQL credentials and database name.

### Database setup

Install [goose](https://github.com/pressly/goose):

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Create the database, then run the migrations from the repo root:

```bash
createdb gator
goose -dir sql/schema postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" up
```

## Commands

### User management

```bash
# Register a new user
gator register <username>

# Log in as an existing user
gator login <username>

# List all users
gator users
```

### Feed management

```bash
# Add a new RSS feed (requires login)
gator addfeed <name> <url>

# List all available feeds
gator feeds

# Follow a feed (requires login)
gator follow <url>

# List feeds you follow (requires login)
gator following

# Unfollow a feed (requires login)
gator unfollow <url>
```

### Aggregation

```bash
# Start fetching feeds continuously (e.g., every 30 seconds)
gator agg 30s
```

This runs in the foreground, scraping each feed on the given interval. Keep it running in a separate terminal while you use the other commands.

### Browsing posts

```bash
# View the latest posts from feeds you follow (requires login)
gator browse

# Limit the number of posts shown
gator browse <limit>
```
