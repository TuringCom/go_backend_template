FROM mysql:5.7

MAINTAINER Celestine Ekoh-Ordan <celestine.e@turing.com>

ENV MYSQL_DATABASE 'turing'
ENV MYSQL_ROOT_PASSWORD 'root'
ENV MYSQL_USER 'turing'
ENV MYSQL_PASSWORD 'turing'
ENV MYSQL_PORT '3306'
ENV MYSQL_HOST 'localhost'

COPY ./db/dump.sql /docker-entrypoint-initdb.d/dump.sql

# install dependencies for golang
RUN apt-get update && \
    apt-get install -y --no-install-recommends --no-install-suggests \
    ca-certificates git curl && \
    mkdir -p /goroot && \
    curl https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

# Set environment variables.
ENV GOROOT /goroot
ENV GOPATH /go
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# install govendor to install packages
RUN go get -u github.com/kardianos/govendor

ENV APP_HOME /go/src/github.com/TuringCom/golang_challenge_template/
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY . $APP_HOME

# install dependencies from vendor.json
RUN govendor sync

EXPOSE 80
EXPOSE 3306

COPY turing-entrypoint.sh /turing-entrypoint.sh

RUN ["chmod", "+x", "/turing-entrypoint.sh"]
ENTRYPOINT ["/turing-entrypoint.sh"]

RUN ls -al
