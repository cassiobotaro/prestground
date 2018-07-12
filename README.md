# Prest POC

Application to test prest features

For the demo we need a postgres instance with a database demo and user postgres.

Using docker you can do:

```
docker run --name=pgdemo --rm -e POSTGRES_USER=postgres -e POSTGRES_DB=demo -p 5432:5432 -d postgres:10
```

and then

`go run migrate up`
