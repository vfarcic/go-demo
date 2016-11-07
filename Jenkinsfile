node("docker") {

  checkout scm

  withEnv([
    "COMPOSE_FILE=docker-compose-test-local.yml",
    "COMPOSE_PROJECT_NAME=go-demo-master"
  ]) {

    stage "Unit"
    sh "docker-compose run --rm unit"
    sh "docker build -t go-demo ."

    stage "Staging"
    try {
      sh "docker-compose up -d staging-dep"
      sh 'docker-compose run --rm staging'
    } catch(e) {
      error "Staging failed"
    } finally {
      sh "docker-compose down"
    }

  }

}