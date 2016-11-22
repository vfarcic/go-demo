def call(project) {
    stage "Unit"
    withEnv([
        "COMPOSE_FILE=docker-compose-test-local.yml",
        "COMPOSE_PROJECT_NAME=${project}-${env.BRANCH_NAME}"
    ]) {
        sh "docker-compose run --rm unit"
        sh "docker build -t go-demo ."
    }
}
