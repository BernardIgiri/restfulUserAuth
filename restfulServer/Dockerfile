FROM ubuntu:16.04

RUN mkdir /opt/app
RUN mkdir /opt/app/log
RUN mkdir /opt/app/config

COPY restfulUserAuth /opt/app/
COPY start.sh /opt/app/

RUN useradd -ms /bin/bash application
RUN chown application:application -R /opt/app/
RUN chmod 0660 -R /opt/app/log
RUN chmod 0400 -R /opt/app/config

ENTRYPOINT ["sh", "/opt/app/start.sh"]
