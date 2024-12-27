# # MySQL の起動を待機する
# until nc -z -v -w30 $DB_HOST $DB_PORT; do
#   echo 'Waiting for MySQL to be ready...'
#   sleep 5
# done
# Create the tables
cd ./db/migrations/ && goose mysql "${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up 
# Start the server (前のコマンドを実行すると止まってしまう。)
cd ../../ && air -c ./.air.toml