#!/bin/sh

BIN_PATH="/usr/local/bin/server"

for var in JWT_PRIVATE_KEY SERVER_HOST SERVER_PORT DB_USER DB_PASSWORD DB_DB DB_HOST ; do
    if ! env | grep -c "$var" > /dev/null; then echo "Required variable $var is unset" 1>&2; exit 1; fi
done

case "$@" in
    web)
        exec $BIN_PATH 
        ;;
    test)
        exec echo "Not implemented yet" 1>&2; exit 255
        ;;
    *)
        echo "Unknown command specified - command was $@" 1>&2
        echo "Available commands: web, test" 1>&2
        exit 1 
        ;;
esac
