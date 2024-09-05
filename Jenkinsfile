pipeline {
    agent any
    tools {
        go 'go1.22'
    }
    stages {
        stage('Checkout') {
            steps {
                // Checkout your code from your version control system
                git 'https://github.com/your-username/your-repo.git'
            }
        }

        stage('Build') {
            steps {
                // Run tests
                sh 'go test ./...'

                // Build the project
                sh 'go build -o myapp'
            }
        }

        stage('Archive') {
            steps {
                // Archive the built artifact
                archiveArtifacts artifacts: 'myapp', fingerprint: true
            }
        }
    }

    post {
        always {
            // Clean up workspace
            cleanWs()
        }
    }
}