DIR="./models"
if [ -d "$DIR" ]; then
  # Take action if $DIR exists. #
  rm -rf $DIR
fi

cd .sqlboiler/ && docker-compose up -d && sleep 5 && make migration-up &&  sqlboiler psql && mv -f models/ ../
docker-compose down -v
