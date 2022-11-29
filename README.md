# shorgot 

URL shortener API written in Golang and PostgreSQL.

# Deploy

[![Deploy on Railway](https://railway.app/button.svg)](https://railway.app/new/template/6JdMnz)


SECRET_KEY is used to create JWT tokens. You can generate one here: https://go.dev/play/p/E8Bm2HdFIzf 

PORT refers to the port where the API will listen. 

DB_CONNECTION is the Postgres URI connection string. You can enter any value and change this variable once you have a running PostgreSQL instance. 

# Routes 

Request examples for each route on Postman. Make sure to update the JWT variable once you log in. 

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/24641176-1b2e360b-dd7e-4faa-992c-b6dd949bab4b?action=collection%2Ffork&collection-url=entityId%3D24641176-1b2e360b-dd7e-4faa-992c-b6dd949bab4b%26entityType%3Dcollection%26workspaceId%3D5d3e6133-e2c4-40aa-a584-4ee628071d0e)
