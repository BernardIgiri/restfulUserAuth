LOCAL_USER_ID := $(shell id -u `whoami`)
build:
	docker-compose build;
	echo "LOCAL_USER_ID="$(LOCAL_USER_ID) > variables.env;
run:
	docker-compose up;
