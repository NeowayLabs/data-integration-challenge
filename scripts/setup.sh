#!/bin/bash
export PG_USER=postgres
export PG_PASS=postgres
export PG_PORT=5432

go get ./...

tar xvzf dataIntegrationChallenge.tgz >> /dev/null

for FILE in $(ls | grep .csv);
do
	echo "Removing header from $FILE";
	vim -u NONE +'1d' +wq! $FILE;
done

IDS=$(docker container ls -q --all --filter 'name=pg')
if [[ $IDS ]]; then docker rm -f $IDS; fi

echo "Creating container 'pg' from $IMG"
docker run --rm \
  --name=pg \
  --hostname=pg \
  --env POSTGRES_USER=$PG_USER \
  --env POSTGRES_PASSWORD=$PG_PASS \
  -p $PG_PORT:5432 \
  -v "$PWD/scripts/schema.sql":/docker-entrypoint-initdb.d/schema.sql \
  -d postgres:alpine \
  -c "log_destination=stderr" \
  -c "log_statement=all"

export PG_HOST=$(docker inspect --format "{{.NetworkSettings.Networks.bridge.IPAddress}}" pg)
