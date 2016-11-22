def call(project) {
    stage "Publish"
    sh "docker tag go-demo localhost:5000/${project}:${env.BRANCH_NAME}.2.${env.BUILD_NUMBER}"
    sh "docker push localhost:5000/${project}:${env.BRANCH_NAME}.2.${env.BUILD_NUMBER}"
}
