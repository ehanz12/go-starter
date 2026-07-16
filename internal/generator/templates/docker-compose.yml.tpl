version: "3.9"

services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: {{PROJECT_NAME}}_app
    restart: unless-stopped
    ports:
      - "${PORT:-8080}:8080"
    env_file:
      - .env
    depends_on:
      db:
        condition: service_healthy
    networks:
      - app_network

  db:
    image: {{DB_IMAGE}}
    container_name: {{PROJECT_NAME}}_db
    restart: unless-stopped
    environment:
      {{DB_ENV}}
    ports:
      - "{{DB_PORT}}:{{DB_PORT}}"
    volumes:
      - db_data:/var/lib/{{DB_DATA_PATH}}
    healthcheck:
      test: {{DB_HEALTHCHECK}}
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - app_network

volumes:
  db_data:

networks:
  app_network:
    driver: bridge
