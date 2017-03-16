#!/usr/bin/env sh

# Test using the Docker database image

if test -z "$DOCKER_IMG"
then
    export DOCKER_IMG="mcilloni/nerdz-test-db"
fi

CONT_NAME="nerdz-test-db"

sudo docker run -d --rm --name "$CONT_NAME" -p 5432:5432 "$DOCKER_IMG" && \
trap "sudo docker stop \"$CONT_NAME\"" INT TERM EXIT && \
go test . 