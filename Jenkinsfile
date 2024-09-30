pipeline {
    agent any

    stages {
        stage('Lint') {
            steps {
                catchError(buildResult: 'Success', stageResult: 'UNSTABLE') {
                    sh {'''
                        golangci-lint -v run
                    '''}
                }
            }
        }

        stage('Build Docker Image') {
            withDockerRegistry(credentialsId: )
        }

        stage('Govulncheck'){
            steps {
                catchError(buildResult: 'Success', stageResult: 'UNSTABLE') {
                    sh {'''
                        govulncheck ./...
                    '''}
                }
            }
        }

    }
}