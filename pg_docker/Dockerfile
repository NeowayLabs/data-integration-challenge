FROM postgres:11-alpine

ENV LANG pt_BR.utf8

ENV POSTGRES_PASSWORD dci
ENV POSTGRES_USER dci
ENV POSTGRES_DB dci

ADD init.sql /docker-entrypoint-initdb.d/.