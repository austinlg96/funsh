sudo docker stop sshd
sudo docker run --name sshd -d --rm -p 2222:2222 funsh 