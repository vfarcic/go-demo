node("docker") {

  checkout scm

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

}