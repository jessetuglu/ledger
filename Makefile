dev:
	docker-compose --file server/docker-compose-dev.yml up --build
prod:
	cd server && docker-compose -f docker-compose-prod.yml up --build
clean:
	cd server && docker-compose -f docker-compose-dev.yml down --volumes