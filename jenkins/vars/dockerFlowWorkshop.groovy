def call(flows) {
    def dockerHost = "tcp://10.100.192.200:2375"
    def proxyHost = "10.100.198.200"
    def proxyReconfPort = 8081
    def consulAddress = "http://10.100.192.200:8500"
    def proxyDockerHost = "tcp://10.100.198.200:2375"
    def registry = "10.100.198.200:5000/"

    withEnv([
            "DOCKER_HOST=${dockerHost}",
            "FLOW_PROXY_HOST=${proxyHost}",
            "FLOW_PROXY_RECONF_PORT=${proxyReconfPort}",
            "FLOW_CONSUL_ADDRESS=${consulAddress}",
            "FLOW_PROXY_DOCKER_HOST=${proxyDockerHost}",
            "REGISTRY=${registry}",
    ]) {
        def args = "--flow=" + flows.join(" --flow=")
        sh "docker-flow ${args}"
    }
}
