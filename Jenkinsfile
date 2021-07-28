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
                                            parameters: [string(name: 'service', value: service)]
                                                pushResult.put(it, pResult.result)
                        }
                    }
                    echo "${pushResult}"
                }
            }
        }
        stage('Deploy services') {
            steps{
                script{
                    testResult.each { service, testStatus ->
                        if (testStatus == 'SUCCESS'){                            
                            sh "make start-dev service=${service}"
                        }
                    }
                }
            }
        }
    }
    post {
        failure {
            sh 'make final-clean-test'           
        }
    }
}