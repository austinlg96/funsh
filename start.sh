sudo docker stop sshd
sudo docker run --name sshd -d --rm -p 22:22 -v /home/ubuntu/funshlogs:/logs/log funsh 