module github.com/kocubinski/cosmos-adr-54/test

go 1.19

replace (
	github.com/kocubinski/cosmos-adr-54/echo/v1/api => ../echo/v1/api/go
	github.com/kocubinski/cosmos-adr-54/echo/v1/foo => ../echo/v1/foo
	github.com/kocubinski/cosmos-adr-54/echo/v2/api => ../echo/v2/api/go
	github.com/kocubinski/cosmos-adr-54/echo/v2/foo => ../echo/v2/foo
)

require (
	github.com/kocubinski/cosmos-adr-54/echo/v1/api v0.0.0-00010101000000-000000000000
	github.com/kocubinski/cosmos-adr-54/echo/v2/api v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.51.0
)

require (
	github.com/cosmos/gogoproto v1.4.3 // indirect
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/kocubinski/cosmos-adr-54/echo/v1/foo v0.0.0-20221130204452-6ff2f16fc6b3 // indirect
	github.com/kocubinski/cosmos-adr-54/echo/v2/foo v0.0.0-00010101000000-000000000000 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
