module github.com/amarnath-ayyadurai-23/microservices

go 1.19

replace github.com/amarnath-ayyadurai-23/microservices/models/ => ../models

require (
	github.com/dimfeld/httptreemux/v5 v5.5.0
	github.com/google/uuid v1.5.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.9
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
