#!/bin/bash

docker run --name rethinkdb \
    -v "$PWD/data:/data" \
    -p 3000:8080 \
    -p 28015:28015 \
    -p 29015:29015 \
    -d docker-rethinkdb
