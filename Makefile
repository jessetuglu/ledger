dev: # starts the dev server
	docker-compose --file server/docker-compose-dev.yml up --build
prod: # starts a prod server
	cd server && docker-compose -f docker-compose-prod.yml up --build
clean: # wipes the db, cleans go mod cache.., etc.
	cd server && docker-compose -f docker-compose-dev.yml down --volumes