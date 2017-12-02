set -e
_term() {
	echo "Caught signal!"
	kill -TERM "$child" 2>/dev/null
}
trap _term TERM INT QUIT KILL STOP

usermod -u $LOCAL_USER_ID application
service nginx start
su - application -c "/opt/app/restfulUserAuth /opt/app/config/config.yaml" &

child=$!
wait "$child"
