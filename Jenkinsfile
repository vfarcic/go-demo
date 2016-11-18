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
      sh "docker-compose run --rm staging"
    } catch(e) {
      error "Staging failed"
    } finally {
      sh "docker-compose down"
    }

    stage "Publish"
    sh "docker tag go-demo localhost:5000/go-demo:2.${env.BUILD_NUMBER}"
    sh "docker push localhost:5000/go-demo:2.${env.BUILD_NUMBER}"

    mail body: "A new go-demo release 2.${env.BUILD_NUMBER} has been published to the registry.", from: "vfarcic@cloudbees.com", subject: "A new go-demo release", to: "viktor@farcic.com"

  }

}