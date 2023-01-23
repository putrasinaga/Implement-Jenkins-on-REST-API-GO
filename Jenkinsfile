pipeline{
    
    agent any 

    stages{
        stage("A"){
            steps{
                echo "========executing A========"
            }
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