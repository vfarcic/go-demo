def call(project) {
    stage "Staging"
    withEnv([
        "COMPOSE_FILE=docker-compose-test-local.yml",
        "COMPOSE_PROJECT_NAME=${project}-${env.BRANCH_NAME}"
    ]) {
        try {
            sh "docker-compose up -d staging-dep"
            sh "docker-compose run --rm staging"
        } catch(e) {
            error "Staging failed"
        } finally {
            sh "docker-compose down"
        }
    }
}
