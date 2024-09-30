run:
	@go run main.go

runs:
	@sh run.sh

# test:
# 	@go test -v ./...
	
# run: build
# 	@./bin/unittest

# migration:
# 	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

# migrate-up:
# 	@go run cmd/migrate/main.go up

# migrate-down:
# 	@go run cmd/migrate/main.go down


# tinggal masukan cmd
# make run
#  make comanda buatan