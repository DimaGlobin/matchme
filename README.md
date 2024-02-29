# MatchMe

## Description.

There is a code of backend of dating service "MatchMe".

## Required
*make*

*go*

*docker*

*migrate*

## Migrate installation (easiest way)

```
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
apt-get update
apt-get install -y migrate
```

## Build and run.

```
make postgresinit
make run
```