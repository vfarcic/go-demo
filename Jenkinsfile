node("cd") {
    checkout scm
    dockerFlow("go-demo", ["deploy", "proxy", "stop-old"])
}
