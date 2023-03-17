gen-crt:
	openssl req -nodes -new -x509 -keyout 3cta.eb.mil.br.key -out 3cta.eb.mil.br.crt -days 720 -config req.cnf
	mv 3cta* srv/ssl/
sync:
	rsync -avuz -e 'ssh -p 2001' --exclude='*.tar*' ../docker ctex@10.13.128.196:/home/ctex/
modd:
	modd -f modd.conf
