set -e
usermod -u $LOCAL_USER_ID application
service nginx start
su - application -c "/opt/app/restfulUserAuth /opt/app/config/config.yaml"
