FROM alpine:3.15

EXPOSE 22/tcp

RUN apk add openssh

RUN mkdir /logs

RUN touch /logs/log

ADD --chown=root:root ./main /bin/funsh

ADD --chown=root:root ./keys/ssh_host_rsa_key /etc/ssh/ssh_host_rsa_key

RUN adduser -D -g "John Doe" -u 4357 -s /bin/funsh jdoe

ADD --chown=jdoe:jdoe ./keys/user_key.pub /home/jdoe/.ssh/authorized_keys

ENTRYPOINT ["/usr/sbin/sshd","-D","-p","2222"]
