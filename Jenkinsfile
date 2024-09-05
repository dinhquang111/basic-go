pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'building...'
            }
        }

        stage('Push') {
            steps {
                echo 'pushing...'
            }
        }
    }

    post {
        always {
            echo 'build done...'
        }
    }
}