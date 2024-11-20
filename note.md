GraphQL

- Developed by Facebook, reduces bandwidth usage when fetching data; the client can request only the fields it needs.
- GraphQL is used in APIs that fetch large amounts of data.
- GraphQL acts as an intermediary between the client and server.

Initial GraphQL project

Run the following command to initial GraphQL project.

- Add gqlgen to tools.go

Run the following command to add gqlgen to tools.go.

```bash
printf '//go:build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go

go mod tidy
```

- Initial gqlgen config and generate models

Run the following command to initial gqlgen config and generate models.

```bash
go run github.com/99designs/gqlgen init

go mod tidy
```

Generate models

Run the following command to generate models from the schema.

```bash
go run github.com/99designs/gqlgen generate
```


