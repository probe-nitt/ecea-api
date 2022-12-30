#!/bin/sh

until nc -z -v -w30 $DB_HOST $DB_PORT; do
   echo "Waiting for database connection at $DB_PORT"
   sleep 5
done

echo -e "\e[34m >>> Starting the server \e[97m"
$1
