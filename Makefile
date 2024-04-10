container:
	docker compose -f dev.docker-compose.yml down && \
	docker compose --project-name portfolio -f dev.docker-compose.yml up -d --build
container-prod:
	docker compose -f prod.docker-compose.yml down && \
	docker compose --project-name portfolio -f prod.docker-compose.yml up -d
templ:
	cd src && templ generate --watch
