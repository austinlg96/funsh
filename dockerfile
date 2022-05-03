FROM alpine:3.15

EXPOSE 22/tcp

RUN apk add openssh

RUN mkdir /logs

RUN touch /logs/log

ADD --chown=root:root ./main /bin/funsh

RUN chmod 111 /bin/funsh

ADD --chown=root:root ./keys/ssh_host_rsa_key /etc/ssh/ssh_host_rsa_key

RUN chmod 600 /etc/ssh/ssh_host_rsa_key

RUN adduser --disabled-password --gecos "John Doe" --uid 4357 --shell /bin/funsh --no-create-home jdoe

RUN sed -i 's/jdoe:!:/jdoe:\*:/g' /etc/shadow

ADD --chown=root:root ./keys/user_key.pub /authorized_keys

RUN chmod 644 /authorized_keys

ADD --chown=root:root ./sshd_config /etc/ssh/sshd_config 

ENTRYPOINT ["/usr/sbin/sshd","-D","-p","22"]
