pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'basic-go'
        DOCKER_TAG = "${env.BUILD_NUMBER}"
        DOCKER_CONTAINER_NAME = 'basic-go'
        TELEGRAM_BOT_TOKEN = '6654034396:AAEc3hoa3r11NRfMb9ALhXmjWjNzOvozEds'
        TELEGRAM_CHAT_ID = '5131367719'
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
           script {
                def status = currentBuild.result ?: 'SUCCESS'
                def message = "Build ${env.BUILD_NUMBER} - ${status}\n\nJob: ${env.JOB_NAME}\nDuration: ${currentBuild.durationString}"
                
                sh """
                    curl -s -X POST https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/sendMessage \
                    -d chat_id=${TELEGRAM_CHAT_ID} \
                    -d text="${message}" \
                    -d parse_mode=HTML
                """
            }
            cleanWs()
        }
    }
}