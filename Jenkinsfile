pipeline {
    agent {
        docker {
            image 'golang:1.17'
        }
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                sh 'echo "Tests passed!"'
            }
        }
        stage('Deploy') {
            steps {
                sh 'echo "Deployed!"'
            }
        }
    }
}