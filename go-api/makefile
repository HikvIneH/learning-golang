GCC=go
GCMD=run
GPATH=main.go

run:
	make build
	$(GCC) $(GCMD) $(GPATH)

build:
	make build_db

build_db:
	rm config/db/db_structs.go
	go run config/db/main.go -json=./config/db/config.json
	mv db_structs.go config/db/