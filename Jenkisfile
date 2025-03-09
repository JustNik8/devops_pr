pipeline {
    agent { 
        docker { 
            image 'golang:1.22-alpine' 
        } 
    }

    triggers {
        githubPush()
    }
    
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}