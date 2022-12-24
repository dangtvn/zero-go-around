## Project Struct
```api
.
├── README.md
├── api
│   ├── Dockerfile
│   ├── doc
│   │   ├── main.api
│   │   ├── posts.api
│   │   └── user.api
│   ├── etc
│   │   └── main.yaml
│   ├── internal
│   │   ├── config
│   │   │   └── config.go
│   │   ├── handler
│   │   │   ├── posts
│   │   │   │   └── getpostshandler.go
│   │   │   ├── routes.go
│   │   │   └── user
│   │   │       ├── createuserhandler.go
│   │   │       └── getuserhandler.go
│   │   ├── logic
│   │   │   ├── posts
│   │   │   │   └── getpostslogic.go
│   │   │   └── user
│   │   │       ├── createuserlogic.go
│   │   │       └── getuserlogic.go
│   │   ├── svc
│   │   │   └── servicecontext.go
│   │   └── types
│   │       └── types.go
│   └── main.go
├── deployment-shell.sh
├── go.mod
├── go.sum
└── rpc
    ├── model
    │   └── user
    │       ├── usersmodel.go
    │       ├── usersmodel_gen.go
    │       └── vars.go
    ├── service
    │   └── user
    │       ├── Dockerfile
    │       ├── etc
    │       │   └── user.yaml
    │       ├── internal
    │       │   ├── config
    │       │   │   └── config.go
    │       │   ├── logic
    │       │   │   └── createuserlogic.go
    │       │   ├── server
    │       │   │   └── userserver.go
    │       │   └── svc
    │       │       └── servicecontext.go
    │       ├── types
    │       ├── user
    │       │   ├── user.pb.go
    │       │   └── user_grpc.pb.go
    │       ├── user.go
    │       ├── user.proto
    │       └── userclient
    │           └── user.go
    └── sql
        └── user.sql
```
```
- goctl model mysql ddl -src rpc/sql/user.sql -dir ./rpc/model/user -c
- sh deployment-shell.sh
```
