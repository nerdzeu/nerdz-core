#!/usr/bin/env sh

# Test using the Docker database image

if test -z "$DOCKER_IMG"
then
    export DOCKER_IMG="mcilloni/nerdz-test-db"
fi

CONT_NAME="nerdz-test-db"

export NERDZ_DB_USER="test_db"
export NERDZ_DB_NAME="test_db"

echo -n "Starting Docker container $CONT_NAME: " && \
sudo docker run -d --rm --name "$CONT_NAME" -p 5432:5432 "$DOCKER_IMG" && \
trap "echo -n 'Destroying Docker container: ' && sudo docker stop \"$CONT_NAME\"" INT TERM EXIT && \
echo 'Letting PostgreSQL a few seconds to startup...' && \
sleep 5 && \
echo "Launching tests" && \
go test "$@"