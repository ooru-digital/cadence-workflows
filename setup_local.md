## Setting Cadence locally
[Ref](https://cadenceworkflow.io/docs/get-started/installation/#run-cadence-server-using-docker-compose)
```
docker-compose up
```

### Register a [domain](https://cadenceworkflow.io/GLOSSARY.html)

```
docker run --network=host --rm ubercadence/cli:master --do cadence-samples domain register -rd 1
```