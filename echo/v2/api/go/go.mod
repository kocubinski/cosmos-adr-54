module github.com/kocubinski/cosmos-adr-54/echo/v2/api

go 1.19

require (
	github.com/cosmos/gogoproto v1.4.3
	github.com/kocubinski/cosmos-adr-54/echo/v1/foo v0.0.0-00010101000000-000000000000
	github.com/kocubinski/cosmos-adr-54/echo/v2/foo v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.51.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace (
	github.com/kocubinski/cosmos-adr-54/echo/v1/foo => ./../../../v1/foo
	github.com/kocubinski/cosmos-adr-54/echo/v2/foo => ./../../../v2/foo
)
