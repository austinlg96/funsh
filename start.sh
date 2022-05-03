sudo docker stop sshd
sudo docker run --name sshd -d --rm -p 22 -v ./funshlogs:/logs/log funsh 