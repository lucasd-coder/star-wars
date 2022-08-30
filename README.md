# StarWars

- Golang
- Gin Gonic
- Redis
- MongoDB

## Funcionalidades

- Adicionar um planeta (com nome, clima e terreno)
- Listar planetas
- Buscar por ID
- Buscar por nome
- Remover planeta

### Como executar o projeto

Definir uma arquivo ```.env``` na raiz do projeto com algumas variáveis ambiente exemplo e rodar o arquivo docker-compose.

| NAME                 | VALUE                 |
|----------------------|-----------------------|
| REDIS_PASSWORD       | redis                 |
| REDIS_URL            | host.docker.internal  |
| REDIS_PORT           | 6379                  |
| REDIS_DB             | 4                     |
| GO_PROFILE           | dev                   |
| DATABASE_MONGODB     | starwars              |
| HOST_MONGODB         | host.docker.internal  |
| PORT_MONGODB         | 27017                 |
| LOG_LEVEL            | info                  |
| HTTP_PORT            | 8080                  |
| SWAPI_URL            | https://swapi.dev     |

### CMD docker-compose gerar a imagem
```
docker-compose up -d
```
### CMD docker-compose ver logs da aplicação
```
docker logs starwars -f
```

### Integrações

- *SWAPI - Star Wars API Integrations*: Realizar a busca pela quantidade de aparições em filmes

##### Curl Adicionar um planeta (com nome, clima e terreno)
```
curl --request POST \
  --url http://localhost:8080/star-wars/planets \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Tatooine",
	"climate": "arid",
	"terrain": "desert"
}'
```
#### Curl Listar planetas
```
curl --request GET \
  --url http://localhost:8080/star-wars/planets
```

#### Curl Buscar por ID
```
curl --request GET \
  --url http://localhost:8080/star-wars/planets/key/630d3a9770a2b7e7034a140f
```

#### Curl Buscar por nome
```
curl --request GET \
  --url http://localhost:8080/star-wars/planets/Tatooine
```
#### Curl Remover planeta
```
curl --request DELETE \
  --url http://localhost:8080/star-wars/planets/key/630b7621419f837457644cbb
```
