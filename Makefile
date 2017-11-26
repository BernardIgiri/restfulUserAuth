build-db:
	docker run --name some-mongo -d mongo --auth;
run-db:
	docker exec it some-mongo mongo admin;
