docker stop sn-app-container-spectre
docker stop sn-server-container-spectre

docker rm sn-app-container-spectre
docker rm sn-server-container-spectre

docker rmi sn-app-image-spectre
docker rmi sn-server-image-spectre