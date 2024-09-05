pipeline {
    agent any
    tools {
        go 'go1.22'
    }
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