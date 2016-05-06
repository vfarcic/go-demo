```bash
docker-compose \
    -f docker-compose-test.yml \
    run --rm unit

docker-compose build

docker push vfarcic/go-demo

docker-compose up -d app
```