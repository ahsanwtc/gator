# Gator

A command-line RSS feed aggregator built with Go and PostgreSQL. Gator allows you to manage RSS feeds, follow your favorite sources, and browse the latest posts from the command line.

## Prerequisites

Before installing and running Gator, make sure you have the following installed on your system:

- **Go** (version 1.21 or later) - [Download Go](https://golang.org/dl/)
- **PostgreSQL** - [Download PostgreSQL](https://www.postgresql.org/download/)

## Installation

Install the Gator CLI using Go's built-in package manager:

```bash
go install github.com/ahsanwtc/gator@latest
```

This will install the `gator` binary to your `$GOPATH/bin` directory. Make sure this directory is in your system's PATH.

## Configuration

### 1. Set up PostgreSQL Database

First, create a PostgreSQL database for Gator:

```sql
CREATE DATABASE gator;
```

### 2. Create Configuration File

Gator requires a configuration file located at `~/.gatorconfig.json`. Create this file with your database connection details:

```json
{
  "db_url": "postgres://username:password@localhost/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `username`, `password`, and database connection details with your PostgreSQL credentials.

### 3. Install Goose for Database Migrations

Gator uses Goose for database migrations. Install it using:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### 4. Run Database Migrations

Navigate to the SQL schema directory and run the migrations:

```bash
cd sql/schema
goose postgres "postgres://username:password@localhost:5432/gator?sslmode=disable" up
```

Replace `username`, `password`, and other connection details with your PostgreSQL credentials (same as in your config file).

## Usage

### User Management

**Register a new user:**
```bash
gator register <username>
```

**Login as an existing user:**
```bash
gator login <username>
```

**View all registered users:**
```bash
gator users
```

**Reset all data (removes all users, feeds, and posts):**
```bash
gator reset
```

### Feed Management

**Add a new RSS feed:**
```bash
gator addfeed <feed_name> <feed_url>
```
Example:
```bash
gator addfeed "TechCrunch" "https://techcrunch.com/feed/"
```

**View all available feeds:**
```bash
gator feeds
```

**Follow a feed:**
```bash
gator follow <feed_url>
```

**View feeds you're following:**
```bash
gator following
```

**Unfollow a feed:**
```bash
gator unfollow <feed_url>
```

### Content Browsing

**Browse recent posts from your followed feeds:**
```bash
gator browse [limit]
```
Example:
```bash
gator browse 10  # Show 10 most recent posts
gator browse     # Show default number of posts
```

**Aggregate new posts from all feeds:**
```bash
gator agg <time_between_requests>
```
Example:
```bash
gator agg 1m     # Fetch new posts every minute
gator agg 30s    # Fetch new posts every 30 seconds
```

## Features

- **User Management**: Register and login system with persistent user sessions
- **Feed Management**: Add, follow, and unfollow RSS feeds
- **Content Aggregation**: Automatically fetch and store new posts from followed feeds
- **Content Browsing**: View recent posts from your followed feeds
- **PostgreSQL Storage**: All data is stored in a PostgreSQL database
- **Concurrent Processing**: Efficient feed processing with Go's concurrency features

## Project Structure

```
gator/
├── main.go              # Main application entry point
├── handlers.go          # Command handlers
├── types.go            # Type definitions
├── middlewares.go      # Authentication middleware
├── internal/
│   ├── config/         # Configuration management
│   ├── database/       # Database queries and models
│   ├── rss/           # RSS feed parsing
│   └── services/      # Business logic services
└── sql/
    ├── queries/       # SQL queries
    └── schema/        # Database schema migrations
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
