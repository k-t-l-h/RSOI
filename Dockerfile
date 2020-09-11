FROM golang:1.14 AS builder

ADD . /opt/app
WORKDIR /opt/app
RUN go build ./cmd/main.go

FROM ubuntu:20/04

RUN apt-get - y update && apt-get install -y tzdata

ENV TZ=Russina/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV PostgresVer 12
ENV PostgresPort 5432

RUN apt-get -y update && apt-get install -y postgresql-$PostgresVer

USER postgres

RUN /ect/init.d/postgresql start &&\
psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" &&\
createdb -O docker docker &&\
/ect/init.d/postgresql stop

EXPOSE $PostgresPort

USER root

WORKDIR /usr/src/app

COPY . .
COPY --from=build /opt/app/main .

EXPOSE 5000
ENV PGpassword docker
CMD service postgresql start && psql -h localhost -d docker -U docker -p $PostgresPort -a -q -f ./init.sql && ./main