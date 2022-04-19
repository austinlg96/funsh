go build 

mkdir keys
ssh-keygen -f ./keys/ssh_host_rsa_key -N "" -t rsa
ssh-keygen -f ./keys/user_key -N "" -t rsa

sudo docker build -t funsh .