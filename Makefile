ssh: 
	ssh root@161.35.247.132

scp-frontend-dotenv:
	scp frontend/.env.prd root@161.35.247.132:~/languagequiz/frontend/.env
