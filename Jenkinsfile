def serviceName = "go-demo"

node("cd") {
    checkout scm

    stage "Unit Tests"
    sh "docker-compose -f docker-compose-test.yml run --rm unit"

    stage "Build"
    sh "docker build -t vfarcic/go-demo ."
    // sh "docker push vfarcic/go-demo"

    stage "Deploy"
    dockerFlow(serviceName, ["deploy", "proxy", "stop-old"])

    stage "Production Tests"
    withEnv(["HOST_IP=10.100.198.200"]) {
        sh "docker-compose -f docker-compose-test.yml run --rm production"
    }

    stash includes: 'consul_*.ctmpl', name: 'consul'
}
node("swarm-master") {
    stage "Health"
    unstash "consul"
    sh "sudo consul-template -consul 10.100.192.200:8500 \
        -template 'consul_service.ctmpl:/data/consul/config/${serviceName}.json' \
        -once"
    sh "sudo consul-template -consul 10.100.192.200:8500 \
        -template 'consul_check.ctmpl:/data/consul/config/${serviceName}_check.json:killall -HUP consul' \
        -once"
}
