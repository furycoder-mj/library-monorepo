def testResult = [:]
def pushResult = [:]
pipeline {
    agent any
    environment {
        baseRegistry = "mihiratdocker/library_" 
        registryCredential = 'dockerhub_id' 
        dockerImage = ''
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
                        script: 'git diff --dirstat=files,0 HEAD~1 | sed -E "s/^[ 0-9.]+% //g" | sed -n "/src\\//p" |sed -E "s/src\\///g" | sed -E "s/\\/.*$//g" | sort | uniq | sed \':a;N;$!ba;s/\\n/,/g\' ',
                        returnStdout: true
                    ).trim()
                    CHANGED_SERVICES_LIST = CHANGED_SERVICES_STR.split(',')
                    echo "changes found in services - ${CHANGED_SERVICES_LIST}"
                }
            }
        }
        stage('Testing all services') {
            steps{
                script{
                    CHANGED_SERVICES_LIST.each {
                        def result = build propagate: false, job: 'LibraryTest',
                                        parameters: [string(name: 'service-name', value: it)]
                                            testResult.put(it, result.result)
                    }
                    echo "${testResult}"
                }
            }
        }
        stage('Build and push images') {
            steps{
                script{
                    testResult.each { service, testStatus ->
                        if (testStatus == 'SUCCESS'){
                            def pResult = build propagate: false, job: 'LibraryBuildAndPush',
                                            parameters: [string(name: 'service-name', value: it)]
                                                pushResult.put(it, pResult.result)
                            // step([$class: 'DockerBuilderPublisher', cleanImages: false, 
                            //  cleanupWithJenkinsJobDelete: false, cloud: '',
                            //  dockerFileDirectory: './src/users-service', fromRegistry: [],
                            //  pushCredentialsId: 'dockerhub_id', pushOnSuccess: true, 
                            //  tagsString: 'mihiratdocker/library_users_service:latest'])                      
                            // registry = baseRegistry + service
                            // dockerImage = docker.build registry + ":$BUILD_NUMBER"
                            // docker.withRegistry( '', registryCredential ) {
                            //     dockerImage.push('latest')
                            // } 
                        }
                    }
                    echo "${pushResult}"
                }
            }
        }
        // stage('Deploy services') {
        //     steps{
        //         script{
        //             testResult.each { service, testStatus ->
        //                 if (testStatus == 'SUCCESS'){                            
        //                     // make deploy... 
        //                 }
        //             }
        //         }
        //     }
        // }
    }
    post {
        failure {
            sh 'make final-clean-test'           
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