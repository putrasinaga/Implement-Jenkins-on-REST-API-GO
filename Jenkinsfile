pipeline{
    
    agent any 
    tools {
        go '1.19.4'
    }

    stages{
        stage("Build Go Project"){
            steps{
                echo "========executing========"
                sh'go build'

            }
        }
    }

    post{
    always{
        echo "menjalankan automation"
    }
    success{
        echo "deploy berhasil"
    }
    failure{
        echo "gagal, deploy"
    }
    cleanup{
        echo "telah proses selesai"
    }
}

}

