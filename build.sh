go build 

mkdir keys
ssh-keygen -q -f ./keys/ssh_host_rsa_key -N "" -t rsa
ssh-keygen -q -f ./keys/user_key -N "" -t rsa

sudo docker build -t funsh .