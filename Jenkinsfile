pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'basic-go'
        DOCKER_TAG = "${env.BUILD_NUMBER}"
        DOCKER_CONTAINER_NAME = 'basic-go'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                script {
                    echo 'Building docker image...'
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}", ".")
                    echo 'Running container...'
                    docker.image("${DOCKER_IMAGE}:${DOCKER_TAG}").run("-p 8080:8080 --name ${DOCKER_CONTAINER_NAME}")
                }  
            }
        }
    }

    post {
        always {
            echo 'build done...'
        }
    }
}