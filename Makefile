ssh: 
	ssh root@161.35.247.132

scp-frontend-dotenv:
	scp frontend/.env.prd root@161.35.247.132:~/languagequiz/frontend/.env

scp-nginx-config:
	scp nginx/nginx.conf root@161.35.247.132:/etc/nginx/conf.d/languagequiz.conf
