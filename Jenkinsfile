pipeline {
    agent { 
        docker { 
            image 'golang:1.22-alpine' 
        } 
    }


    environment {
        GO111MODULE = 'on'
    }

    triggers {
        githubPush()
    }
    
    stages {

        stage('Check Go Environment') {
            steps {
                sh 'go version'
            }
        }


        stage('Checkout') {
            steps {
                // Check code from SCM (Git) for all branches
                checkout([$class: 'GitSCM', 
                          branches: [[name: '**']], 
                          userRemoteConfigs: [[
                              url: 'git@github.com:JustNik8/devops_pr.git',
                              credentialsId: 'github-ssh-key'
                          ]]])
            }
        }

        stage('Build') {
            steps {
                script {
                    sh '''
                    go mod tidy
                    go build -o devops_pr
                    '''
                }
            }
        }


        stage('Test') {
            steps {
                script {
                    sh 'go test ./...'
                }
            }
        }
        
    }
}