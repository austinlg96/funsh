go build
cp ./main /bin/funsh
echo "/bin/funsh" >> /etc/shells
mkdir /logs
chown root:root /logs
chmod 001 /logs
touch /logs/log
chmod 002 /logs