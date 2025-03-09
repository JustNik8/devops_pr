pipeline {
    agent { 
        docker { 
            image 'golang:1.22-alpine' 
            args '-v ${WORKSPACE}:/go/src/devops_pr'
        } 
    }


    environment {
        GO111MODULE = 'on'
        GOCACHE = "/go/src/devops_pr/.cache/go-build"
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

        stage('Prepare Environment') {
            steps {
                script {
                    // Создание директории для кэша, если она не существует
                    sh 'mkdir -p /go/src/devops_pr/.cache/go-build'
                }
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
                    go build -v ./...
                    '''
                    // go build -o devops_pr
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