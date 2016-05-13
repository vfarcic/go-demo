node("cd") {
    checkout scm
    dockerFlow("docker-flow", ["deploy", "proxy", "stop-old"])
}
