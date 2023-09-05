#wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"
wait-for db:3306 -t 60

#watch for .go file changes
# Watch your .go files and invoke go build if the files changed.
CompileDaemon --build="go build -o ./cmd/main ./cmd/main.go"  --command=./cmd/main