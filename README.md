# Docker
* Image = molde das dependências do ambiente

docker run <imagem>

## Listar containers
docker ps

## Listar containers historico
docker ps -a

## Sincronizar porta do computador para uma porta do container
docker run -p 8080:80 <image>

## Processo em background (desataxada)
docker run -d -p <image>

## Parar container
docker stop <CONTAINER ID>

## Remover container
docker rm -f <CONTAINER ID>

## Nomear container
docker run --name <NOME> -d -p 8080:80 <image>

docker exec <iamge> ls

## Entrar de forma iterativa no container logar no (container)
docker exer -it <NOME> bash

## Lista de imagens no computador
docker images
* ajuda em cache para imagens futuras

## Remover imagem
docker rmi <image>

## Criar imagem
docker build -t <name-image> <caminho-dockerfile>

## Subir images docker hub
docker push <name-image>

# Docker-compose
Usado para subir vários contêiners

## Subir serviços
docker-compose up
docker-compose up -d

## Derrubar tudo
docker-compose down
