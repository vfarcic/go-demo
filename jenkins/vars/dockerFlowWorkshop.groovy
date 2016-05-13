def call(project, flows) {
    withEnv([
            "DOCKER_HOST=tcp://10.100.192.200:2375",
            "FLOW_PROXY_HOST=10.100.198.200",
            "FLOW_PROXY_RECONF_PORT=8081",
            "FLOW_CONSUL_ADDRESS=http://10.100.192.200:8500",
            "FLOW_PROXY_DOCKER_HOST=tcp://10.100.198.200:2375",
    ]) {
        def args = "-p " + project + " --flow=" + flows.join(" --flow=")
        sh "docker-flow ${args}"
        sh "docker ps -a"
    }
}
