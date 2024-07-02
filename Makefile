.PHONY: all build-backend build-frontend deploy

all: build-frontend build-backend deploy

build-backend:
    go build -o main main.go

build-frontend:
    cd frontend && npm run build && cd ..
    mv frontend/build backend/frontend/build

deploy:
    cd backend && gcloud app deploy