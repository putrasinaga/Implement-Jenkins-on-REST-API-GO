pipeline{
    agent any 
    tools {
        go '1.19'
    }

 parameters{
    string(name : 'NAME', defaultValue: 'Guest', description : 'silahkan masukan nama docker')
    booleanParam(name: 'DEPLOY', defaultValue: false, description: 'apakah ingin dilanjutkan?')
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
                sh"docker build -t putrasaut/${params.Name} ."
        }
    }
        stage("push to DockerHUb"){
            steps{
                echo "========Pushing======"
                script {
                   withCredentials([string(credentialsId: 'dockerhubpassword', variable: 'dockerhubpwd')]){
                     sh 'docker login -u putrasaut -p ${dockerhubpwd}'
                    }
                     sh "docker push putrasaut/${params.Name}"
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



