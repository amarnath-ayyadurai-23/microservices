module github.com/amarnath-ayyadurai-23/microservices

go 1.19

replace github.com/amarnath-ayyadurai-23/microservices/models/ => ../models

require (
	github.com/dimfeld/httptreemux/v5 v5.5.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.9
)
