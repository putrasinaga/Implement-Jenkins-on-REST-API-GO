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

