FROM python:3.6-alpine

COPY buildDB.py /root/
COPY config.yaml /root/
COPY config.yaml.key /root/
COPY requirements.txt /root/

RUN pip3 install -r /root/requirements.txt

ENTRYPOINT ["python3", "/root/buildDB.py"]
