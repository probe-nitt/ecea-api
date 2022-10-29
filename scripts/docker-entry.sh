#!/bin/sh

until nc -z -v -w30 probe_db 5432; do
   echo "Waiting for database connection..."
   sleep 5
done

echo -e "\e[34m >>> Starting the server \e[97m"
$1
