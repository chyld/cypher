# cypher

```
psql -h 0.0.0.0 -U postgres -d temp -f schema.sql
watch "golint .; gofmt -w .; go build"
watch -n 0.5 ./build.sh
```
