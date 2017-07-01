pipeline {
  agent {
    label "docker"
  }
  stages {
    stage("Unit") {
      steps {
        git "https://github.com/vfarcic/go-demo.git"
        // sh "docker-compose -f docker-compose-test.yml run --rm unit"
        sh "docker image build -t go-demo ."
      }
    }
    stage("Staging") {
      steps {
        sh "docker-compose -f docker-compose-test-local.yml up -d staging-dep"
        sh 'HOST_IP=localhost docker-compose -f docker-compose-test-local.yml run --rm staging'
      }
    }
    stage("Publish") {
      steps {
        sh "docker tag go-demo localhost:5000/go-demo"
        sh "docker tag go-demo localhost:5000/go-demo:2.${env.BUILD_NUMBER}"
        sh "docker push localhost:5000/go-demo"
        sh "docker push localhost:5000/go-demo:2.${env.BUILD_NUMBER}"
      }
    }
    stage("Production") {
      steps {
        sh "DOCKER_HOST=tcp://${env.PROD_IP}:2375 docker service update --image localhost:5000/go-demo:2.${env.BUILD_NUMBER} go-demo_main"
        try {
          for (i = 0; i < 10; i++) {
            sh "HOST_IP=${env.PROD_IP} docker-compose -f docker-compose-test-local.yml run --rm production"
          }
        } catch (e) {
          sh "DOCKER_HOST=tcp://${env.PROD_IP}:2375 docker service update --rollback go-demo"
        }
      }
    }
  }
  post {
    always {
      sh "docker-compose -f docker-compose-test-local.yml down"
    }
  }
}
