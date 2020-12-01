# hello-world-go-pgx

This repo has a "Hello World" Go application that uses the [pgx driver](https://pkg.go.dev/github.com/jackc/pgx) to talk to [CockroachDB](https://www.cockroachlabs.com/docs/stable/).

To run the code:

1. Start a [demo CockroachDB cluster](https://www.cockroachlabs.com/docs/stable/cockroach-demo.html) from the command line: `cockroach demo --empty`

1. Create a `bank` database and a user and password as described in [Build a Go app with CockroachDB](https://www.cockroachlabs.com/docs/stable/build-a-go-app-with-cockroachdb.html).

1. From the SQL client: `GRANT ALL ON DATABASE bank TO user`

1. Update the connection string passed to `pgx.ParseConfig()` with the correct connection values.

1. In your terminal, from this directory: `go mod init basic-sample && go run main.go`
