## Start the server
```sh
make pre-commit
go run .
```

## Migrations
```sh
# Create migration file
make migrate-create filename=namefile

# Running up migrations
make migration-up

# Running down migrations
make migration-down
```


## Generate Model
```sh
make gen-model
```

## Install SQLBoiler
```sh
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```