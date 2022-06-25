dev:
	docker-compose --file server/docker-compose-dev.yml --build
prod:
	cd server && docker-compose -f docker-compose-prod.yml --build
clean:
	cd server && docker-compose -f docker-compose-dev.yml --build