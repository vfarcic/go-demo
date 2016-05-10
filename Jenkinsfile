node("cd") {

export DOCKER_HOST=tcp://10.100.192.200:2375
export FLOW_PROXY_HOST=10.100.198.200
export FLOW_PROXY_RECONF_PORT=8081
export FLOW_CONSUL_ADDRESS=http://10.100.192.200:8500
export FLOW_PROXY_DOCKER_HOST=tcp://10.100.198.200:2375
export REGISTRY=10.100.198.200:5000/

docker-flow --flow=deploy --flow=proxy --flow=stop-old

}