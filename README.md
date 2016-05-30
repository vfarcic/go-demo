```bash
scripts/setup.sh

docker-compose \
    -f docker-compose-test.yml \
    run --rm unit

docker build -t vfarcic/go-demo .

docker tag vfarcic/go-demo vfarcic/go-demo:1.0

docker tag vfarcic/go-demo vfarcic/go-demo:1.1

docker push vfarcic/go-demo

docker push vfarcic/go-demo:1.0

docker push vfarcic/go-demo:1.1

docker-compose up -d db app
```


```bash
docker-machine create -d virtualbox jenkins

eval $(docker-machine env jenkins)

open http://$(docker-machine ip jenkins):8080

# Install "Pipeline" Plugin
# Install "Docker Slaves" Plugin





docker run -v /var/run/docker.sock:/var/run/docker.sock -ti docker
```