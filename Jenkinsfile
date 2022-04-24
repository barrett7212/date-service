pipeline {
    agent any
    tools {
        // See https://levelup.gitconnected.com/automating-build-and-test-of-go-application-with-jenkins-9f96879b5479
        go 'Go 1.17'
    }
     environment {
        GO117MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Build') {
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