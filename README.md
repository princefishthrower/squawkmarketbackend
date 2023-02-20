# squawkmarketbackend

The go based backend for the squawkmarket app. A sub-second latency stock market squawk box.

## Migrations

Migrations are managed by [golang-migrate](https://github.com/golang-migrate/migrate)

### Create a New Migration

```bash
/bin/bash dev/draft-migration.sh my_example_migration
```

This will scaffold both an `up` `down` migration file in the `db/migrations` folder.

### Run Database Migrations

```bash
/bin/bash ./dev/run-migrations.sh
```

## Run Down Migrations

**WARNING THIS CAN CAUSE DATA LOSS! RUN AT YOUR OWN RISK! YOU HAVE BEEN WARNED!**

```bash
/bin/bash ./dev/run-migrations-down.sh
```

## Build & Run Docker Locally

```shell
source .env.sh && docker build -t squawkmarketbackend:latest .
```

Then to run it:

```shell
docker run --name squawkmarketbackend -d -p 8080:8080 squawkmarketbackend:latest
```

## Build Docker for Managed Server / Cloud

The build is managed by Circle CI, triggered by pushing to the `master` branch. Essentially the process is the same as the above, it logs into the server of choice and issues the above commands.
