FROM mysql:5.7.22

COPY ./docker/custom.config /etc/mysql/conf.d/custom.config