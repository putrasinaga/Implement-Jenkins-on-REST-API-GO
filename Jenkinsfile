pipeline{
    agent any 
    tools {
        go '1.19'
    }
    stages{
        stage("Build Go Project"){
            steps{
                echo "========executing========"
                sh'go build'
            }
        }
        stage("build container"){
            steps{
                echo "========Build image======"
                sh'docker build -t putrasaut/web-simple-api .'
        }
    }
        stage{"push to DockerHUb"}{
            steps{
                echo "========Pushing======"`
                script {
                    withCredentials([string(credentialsId: 'Dockerhub-pass', variable: 'Dockerhub-pass')]) {
                     sh 'docker login -u putrasaut -p ${Dockerhub-pass}'
                    }
                     sh 'docker push putrasaut/web-simple-api'
                }
            }
        }    
    }
    
    post{
    always{
        echo "menjalankan automation"
    }
    success{
        echo "berhasil"
    }
    failure{
        echo "gagal"
    }
    cleanup{
        echo "telah proses selesai"
    }
}
}



