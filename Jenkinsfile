node("cd") {
    checkout scm
    dockerFlow(["deploy", "proxy", "stop-old"])
}
