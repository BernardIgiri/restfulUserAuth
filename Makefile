LOCAL_USER_ID := $(shell id -u `whoami`)
build:
	echo "LOCAL_USER_ID="$(LOCAL_USER_ID) > variables.env;
	docker-compose build;
run:
	docker-compose up;
