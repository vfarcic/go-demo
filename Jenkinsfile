node("docker") {

  checkout scm

  withEnv([
    "COMPOSE_FILE=docker-compose-test-local.yml",
    "COMPOSE_PROJECT_NAME=go-demo-artifactory"
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

    stage "Artifactory"
    def server = Artifactory.server('artifactory-1')
    def artSpec = '{ "files": [ { "pattern": "go-demo", "target": "go-demo/" } ] }'
    server.upload(artSpec)

    stage "Publish"
    sh "docker tag go-demo localhost:5000/go-demo:2.${env.BUILD_NUMBER}"
    sh "docker push localhost:5000/go-demo:2.${env.BUILD_NUMBER}"

  }

}