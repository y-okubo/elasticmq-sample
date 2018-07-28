## Start ElasticMQ

```
$ docker run -p 9324:9324 s12v/elasticmq
```

## Put the item

```
$ ruby client.rb
```

## Pop the item

```
$ go run wworker.go
```
