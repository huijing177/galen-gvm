.PHONY: swag.install swag run

swag.install:
	if ! which swag >/dev/null 2>&1  ;then \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi;

swag: swag.install
	swag init

run: swag
	go run .