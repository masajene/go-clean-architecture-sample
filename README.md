# Go CleanArchitecture Boilerplate

![UML](doc/app.png)

```
├── Makefile
├── application
│   ├── model
│   │   └── user_model.go
│   ├── repository
│   │   └── user_repository.go
│   └── usecase
│       └── user_usecase.go
├── config
│   └── config.go
├── di
│   ├── wire.go
│   └── wire_gen.go
├── domain
│   ├── entity
│   │   └── user.go
│   ├── interactor
│   │   └── user_interactor.go
│   └── validation
│       └── validation.go
├── infra
│   ├── aws
│   │   ├── aws.go
│   │   ├── s3.go
│   │   ├── ses.go
│   │   ├── sqs.go
│   │   └── ssm.go
│   ├── database
│   │   ├── database.go
│   │   └── mysql.go
│   └── router
│       └── router.go
├── interfaces
│   ├── gateway
│   │   └── user_repository.go
│   └── handler
│       └── user_handler.go
├── lambda.go
├── main.go
├── tmp
│   ├── build-errors.log
│   └── main
└── utils
    ├── error_model.go
    ├── log.go
    └── time.go
```