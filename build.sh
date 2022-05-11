go build 

mkdir keys
ssh-keygen -q -f ./keys/ssh_host_rsa_key -N "" -t rsa
ssh-keygen -q -f ./keys/user_key -N "" -t rsa

sudo docker build -t funsh .
sudo docker save -o "funsh_image.tar" funsh
sudo chmod 444 funsh_image.tar