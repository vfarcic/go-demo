def call(project) {
    stage "Production"
    withEnv([
        "COMPOSE_FILE=docker-compose-test-local.yml",
        "COMPOSE_PROJECT_NAME=${project}-${env.BRANCH_NAME}",
        "DOCKER_TLS_VERIFY=1",
        "DOCKER_HOST=tcp://${env.PROD_IP}:2376",
        "DOCKER_CERT_PATH=/machines/${env.PROD_NAME}"
    ]) {
        sh "docker service update --image localhost:5000/${project}:${env.BRANCH_NAME}.2.${env.BUILD_NUMBER} ${project}"
        for (i = 0; i < 10; i++) {
            sh "HOST_IP=${env.PROD_IP} docker-compose run --rm production"
        }
    }
}
