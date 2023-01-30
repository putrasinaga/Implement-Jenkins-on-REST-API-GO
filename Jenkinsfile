pipeline{
    agent any 
    tools {
        go '1.19'
    }

 parameters{
    string(name : 'NAME', defaultValue: 'nama-project', description : 'silahkan masukan nama docker')
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
                sh"docker build -t putrasaut/${params.NAME} ."
        }
    }
        stage("push to DockerHUb"){
            input{
                message "lanjutkan deploy?"
                ok "iya"
                submitter "saut"
            }
            steps{
                echo "========Pushing======"
                script {
                   withCredentials([string(credentialsId: 'dockerhubpassword', variable: 'dockerhubpwd')]){
                     sh 'docker login -u putrasaut -p ${dockerhubpwd}'
                    }
                     sh "docker push putrasaut/${params.NAME}"
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



