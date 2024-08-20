docker system prune

docker build -t dockerize .

docker run -p 8081:8081 dockerize