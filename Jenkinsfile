pipeline {
    agent any
    // tools {
    //     go 'go1.16.5'
    // }
    environment {
        // GO114MODULE = 'on'
        // CGO_ENABLED = 0 
        // GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        // registry = "mihiratdocker/jenkins_golang_hello_world_pipeline" 
        // registryCredential = 'dockerhub_id' 
        dockerImage = ''
        CHANGED_SERVICES_STR = ''
        // CHANGED_SERVICES_LIST = []
    }
    stages {      
        stage('Cloning git repo') {
            steps{
                git branch: 'main', credentialsId: 'personal_git_ssh', url: 'git@github.com:furycoder-mj/library-monorepo.git'
            }
        }
        stage('Finding changed services'){
            steps{
                script{
                    CHANGED_SERVICES_STR = sh (
                        script: 'git diff --dirstat=files,0 HEAD~1 | sed -E "s/^[ 0-9.]+% //g" | sed -n "/src\\//p" |sed -E "s/src\\///g" | sed -E "s/\\/.*$//g" | sed -e ":a" -e "N" -e "$!ba" -e "s/\\n/,/g" | tr "-" "_" ',
                        returnStdout: true
                    ).trim()
                    // CHANGED_SERVICES_LIST = CHANGED_SERVICES_STR.split(' ')
                    // echo "changes found in services - ${CHANGED_SERVICES_LIST}"
                }
            }
        }
        stage('Testing all services') {
            //run docker compose for compose.test file
            // make test
            steps{
                script{
                    CHANGED_SERVICES_STR.tokenize(',').each {                         
                            echo "Item: ${it}"    
                            sh "make test service=${it}"
                        }
                }
                // script{
                //     dockerImage = docker.build registry + ":$BUILD_NUMBER" 
                // }
            }
        }
    }
    post {
        failure {
            sh 'make clean-test'           
        }
    }
        // stage('Building our image') {
        //     steps{
        //         script{
        //             dockerImage = docker.build registry + ":$BUILD_NUMBER" 
        //         }
        //     }
        // }
        // stage('Push our image') {
        //     steps{
        //         script{
        //             docker.withRegistry( '', registryCredential ) {
        //                 dockerImage.push('latest')
        //             } 
        //         }
        //     }
        // }
        // stage('Deploy our image'){
        //     steps{
        //         sh 'docker container rm -f testDeployment'
        //         script{
        //             containerId = docker.image('mihiratdocker/jenkins_golang_hello_world_pipeline:6').run('-p 8001:8001 --name testDeployment')
        //         }
        //     }  
        // }
        // stage('Pre Test') {
        //     steps {
        //         echo 'Installing dependencies'
        //         sh 'go version'
        //         sh 'go get -u golang.org/x/lint/golint'
        //     }
        // }
        
        // stage('Build') {
        //     steps {
        //         echo 'Compiling and building'
        //         sh 'go build'
        //     }
        // }

        // stage('Test') {
        //     steps {
        //         withEnv(["PATH+GO=${GOPATH}/bin"]){
        //             echo 'Running vetting'
        //             sh 'go vet .'
        //             echo 'Running linting'
        //             sh 'golint .'
        //             echo 'Running test'
        //             sh 'cd test && go test -v'
        //         }
        //     }
        // }
    // }  
}