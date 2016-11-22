def call(project) {
    withEnv([
            "COMPOSE_FILE=docker-compose-test-local.yml",
            "COMPOSE_PROJECT_NAME=${project}"
    ]) {

        stage "Unit"
        sh "docker-compose run --rm unit"
        sh "docker build -t go-demo ."
    }
}
