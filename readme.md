## Clean Architecture Go Lang

API Rest service using designer pattern with golang as Clean Architecture, SOLID and DDD.

> #### Requirements

- docker
- docker-compose

> #### Used Libraries

- github.com/bxcodec/faker/v3
- github.com/gin-gonic/gin
- github.com/go-playground/validator/v10
- github.com/stretchr/testify
- go.mongodb.org/mongo-driver
- golang.org/x/crypto

> #### Run Tests

- Run command `go test -cover ./...`
- Run command `go test -cover -coverprofile=c.out ./...`
- Run command `go tool cover -html=c.out -o coverage.html`

> #### Start Project

- Run command `docker-compose up -d`

> #### Close Project

- Run command `docker-compose down`
