FROM gomicro/goose

WORKDIR /migrations

ADD /pkg/sql/schema/*.sql ./
ADD /pkg/sql/schema/goose_docker_config/migrations.sh ./migrations.sh

RUN chmod +x ./migrations.sh