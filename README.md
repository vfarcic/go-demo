```bash
docker-compose \
    -f docker-compose-test.yml \
    run --rm unit

docker build -t vfarcic/go-demo .

docker push vfarcic/go-demo

docker-compose up -d
```