
pipeline {
  agent any
  
  stages {
    stage('Notify Build Start') {
      steps {
        slackSend channel: "#engineering", message: "Build Started: ${currentBuild.fullDisplayName} (<${env.RUN_DISPLAY_URL}|BlueOcean> <${env.BUILD_URL}|Open>)"
      }
    }
    stage('Checkout') {
      steps {
        checkout scm
      }
    }
    stage('Prepare') {
      steps {
        sh 'make clean'
        sh 'make bootstrap'
      }
    }
    stage('Build') {
      steps {
        sh 'make'
      }
    }
  }

  post {
    success {
      slackSend color: "good", channel: "#engineering", message: "Build finished successfully: ${currentBuild.fullDisplayName} (<${env.RUN_DISPLAY_URL}|BlueOcean> <${env.BUILD_URL}|Open>)"
    }
    failure {
      slackSend color: "danger", channel: "#engineering", message: "Build failed: ${currentBuild.fullDisplayName} (<${env.RUN_DISPLAY_URL}|BlueOcean> <${env.BUILD_URL}|Open>)"
    }
  }
}