def serviceName = "go-demo"

node("cd") {
    checkout scm

    stage "Deploy"
    dockerFlow(serviceName, ["deploy", "proxy", "stop-old"])
    stash includes: 'consul_*.ctmpl', name: 'consul'
}
node("swarm-master") {
    stage "Health"
    unstash "consul"
    sh "sudo consul-template -consul localhost:8500 \
        -template 'consul_service.ctmpl:/data/consul/config/${serviceName}.json' \
        -once"
    sh "sudo consul-template -consul localhost:8500 \
        -template 'consul_check.ctmpl:/data/consul/config/${serviceName}_check.json:killall -HUP consul' \
        -once"
}
