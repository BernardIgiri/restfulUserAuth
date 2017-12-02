set -e
usermod -u $LOCAL_USER_ID application
su - application -c "/opt/app/restfulUserAuth /opt/app/config/config.yaml"
