FROM ubuntu/latest

COPY restfulUserAuth /usr/src/app/

RUN useradd -ms /bin/bash application
RUN chown application:application -R /user/src/app/
RUN chmod 0660 -R /user/src/app/log
RUN chmod 0400 -R /user/src/app/config

CMD ["sudo", "-u", "restfulUserAuth", "config/config.yaml"]
