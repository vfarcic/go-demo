node("docker") {

  checkout scm

  withEnv([
    "COMPOSE_FILE=docker-compose-test-local.yml",
    "COMPOSE_PROJECT_NAME=go-demo-trust-issues"
  ]) {

    stage "Unit"
    sh "docker-compose run --rm unit"
    sh "docker build -t go-demo ."

    stage "Staging"
    try {
      sh "docker-compose up -d staging-dep"
      sh "docker-compose run --rm staging"
    } catch(e) {
      error "Staging failed"
    } finally {
      sh "docker-compose down"
    }

    stage "Publish"
    sh "docker tag go-demo localhost:5000/go-demo:2.${env.BUILD_NUMBER}"
    sh "docker push localhost:5000/go-demo:2.${env.BUILD_NUMBER}"

    stash "docker-compose-test-local.yml"
  }
}

input "Proceed with deployment?"

input message: 'Do you really want to proceed with deployment?', submitter: 'admin'

node("docker") {
  unstash "docker-compose-test-local.yml"

  stage "Production"
  withEnv([
    "DOCKER_TLS_VERIFY=1",
    "DOCKER_HOST=tcp://${env.PROD_IP}:2376",
    "DOCKER_CERT_PATH=/machines/${env.PROD_NAME}"
  ]) {
    sh "docker service update --image localhost:5000/go-demo:2.${env.BUILD_NUMBER} go-demo"
  }
    withEnv([
      "COMPOSE_FILE=docker-compose-test-local.yml",
      "COMPOSE_PROJECT_NAME=go-demo-trust-issues"
    ]) {
    for (i = 0; i < 10; i++) {
      sh "HOST_IP=${env.PROD_IP} docker-compose run --rm production"
    }
  }

}