FROM ubuntu
RUN apt-get update && apt-get install -y \
    build-essential \
    vim \
    wget

RUN wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz

RUN mkdir -p /root/bin \
 && mkdir -p /root/pkg

WORKDIR /root
ADD . /root/

ENV GOPATH="/root"
ENV PATH="${PATH}:/usr/local/go/bin:${GOPATH}/bin"

RUN go install github.com/tony-yang/gear-designer
