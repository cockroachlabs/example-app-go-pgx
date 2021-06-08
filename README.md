# hello-world-go-pgx

This repo has a "Hello World" Go application that uses the [pgx driver](https://pkg.go.dev/github.com/jackc/pgx) to talk to [CockroachDB](https://www.cockroachlabs.com/docs/stable/).

To run the code:

1. Start a [demo CockroachDB cluster](https://www.cockroachlabs.com/docs/stable/cockroach-demo.html) from the command line: `cockroach demo --empty`

1. Create a `bank` database and a user and password as described in [Build a Go app with CockroachDB](https://www.cockroachlabs.com/docs/stable/build-a-go-app-with-cockroachdb.html).

1. From the SQL client: `GRANT ALL ON DATABASE bank TO demo`

1. Update the `{demo_password}` in the connection string passed to `pgx.ParseConfig()` with the correct password. `cockroach demo` generates a new password for each new virtual cluster.

1. In your terminal, from this directory: `go mod init basic-sample && go mod tidy && go run main.go`
