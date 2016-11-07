node("docker") {

  git "https://github.com/vfarcic/go-demo.git"

  stage "Unit"
  sh "docker-compose -f docker-compose-test.yml run --rm unit"
  sh "docker build -t go-demo ."

  stage "Staging"
  try {
    sh "docker-compose -f docker-compose-test-local.yml up -d staging-dep"
    sh 'HOST_IP=localhost docker-compose -f docker-compose-test-local.yml run --rm staging'
  } catch(e) {
    error "Staging failed"
  } finally {
    sh "docker-compose -f docker-compose-test-local.yml down"
  }

  stage "Publish"
  sh "docker tag go-demo localhost:5000/go-demo:2.${env.BUILD_NUMBER}"
  sh "docker push localhost:5000/go-demo:2.${env.BUILD_NUMBER}"

  stage "Production"
  withEnv([
    "DOCKER_TLS_VERIFY=1",
    "DOCKER_HOST=tcp://${env.PROD_IP}:2376",
    "DOCKER_CERT_PATH=/machines/${env.PROD_NAME}"
  ]) {
    sh "docker service update --image localhost:5000/go-demo:2.${env.BUILD_NUMBER} go-demo"
  }
  sh "HOST_IP=${env.PROD_IP} docker-compose -f docker-compose-test-local.yml run --rm production"

}