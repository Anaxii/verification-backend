git pull
sudo docker build -t kyc-backend .
sudo docker container stop container `sudo docker ps -aq`
sudo docker run -p 80:8080 kyc-backend:latest