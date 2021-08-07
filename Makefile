runbackend:
	docker-compose up -d && go run ./server
runweb:
	npm run --prefix ./app dev
rundb:
	docker-compose up
stopdb:
	docker-compose down
setEnv:
	cp .env.example .env