# library-monorepo
This repo demonstrates use of microservices monorepo concept.  
*Monorepo* is attributed to multiple microservices of a system residing in single repository.  
  
Jenkins is used for CICD pipeline.  
Separate JenkinsFile have been used for different pipelines-
1. Jenkinsfile.LibraryTest   
2. Jenkinsfile.LibraryBuildPush  
3. Jenkinsfile.LibraryDeploy  
A main pipeline detects changes in git from last diff and triggers these secondary pipelines as needed.

Makefiles are used to run sequence of commands declaratively from Jenkinsfile.  
For encapsulating the complex commands and separation of concern.  
  
Docker compose is used to deploy these services locally.  

A webhook is configured from github to jenkins.
A tunneling service called *ngrok* is used for testing purpose and basically exposes the jenkins running on localhost to internet.

