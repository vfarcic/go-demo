def call(project, flows, args = []) {
    withEnv([
            "DOCKER_HOST=tcp://10.100.192.200:2375",
            "FLOW_PROXY_HOST=10.100.198.200",
            "FLOW_PROXY_RECONF_PORT=8081",
            "FLOW_CONSUL_ADDRESS=http://10.100.192.200:8500",
            "FLOW_PROXY_DOCKER_HOST=tcp://10.100.198.200:2375",
    ]) {
        def dfArgs = "-p " + project + " --flow=" + flows.join(" --flow=") + " " + args.join(" ")
        sh "docker-flow ${dfArgs}"
        sh "docker ps -a"
    }
}
