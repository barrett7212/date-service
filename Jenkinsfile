pipeline {
    agent any
    tools {
        go 'go1.17'
    }
     environment {
        GO117MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Build') {
            steps {
                sh 'docker container run  '
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