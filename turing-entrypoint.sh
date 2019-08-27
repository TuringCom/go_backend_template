#!/bin/bash
set -x;

/bin/bash /entrypoint.sh mysqld > /dev/null 2>&1 &

# wait for mysql
echo "Waiting for MySQL"
until echo '\q' | mysql -h "$MYSQL_HOST" -P "$MYSQL_PORT" -u root -p"$MYSQL_ROOT_PASSWORD" $MYSQL_DATABASE; do
    >&2 echo "MySQL is unavailable - Sleeping..."
    sleep 2
done

echo -e "\nMySQL ready!"

# set port for autohost
export PORT=80

go run main.go
