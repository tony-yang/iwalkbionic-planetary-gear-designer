FROM ubuntu
RUN apt-get update && apt-get install -y \
    build-essential \
    vim \
    wget

RUN wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz

ENV GOPATH="${HOME}"
ENV PATH="${PATH}:/usr/local/go/bin:${GOPATH}/bin"

WORKDIR /root

ADD . /root/
