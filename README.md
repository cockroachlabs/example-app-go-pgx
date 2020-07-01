# hello-world-go-pgx

This repo has a "Hello World" Go application that uses the [pgx driver](https://pkg.go.dev/github.com/jackc/pgx) to talk to [CockroachDB](https://www.cockroachlabs.com/docs/stable/).

To run the code:

1. Start a [local, insecure CockroachDB cluster](https://www.cockroachlabs.com/docs/stable/start-a-local-cluster.html).

2. Create a `bank` database and `maxroach` user as described in [Build a Go app with CockroachDB](https://www.cockroachlabs.com/docs/stable/build-a-go-app-with-cockroachdb.html#insecure).

3. From the [SQL client](https://www.cockroachlabs.com/docs/stable/cockroach-sql.html): `GRANT ALL ON DATABASE bank TO maxroach`

4. In your terminal, from this directory: `go mod init basic-sample && go run main.go`
