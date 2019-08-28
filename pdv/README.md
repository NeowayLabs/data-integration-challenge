
## Execução

Para execução do projeto execute:

```
  docker-compose up --build
```

## Estrutura
### pdv-database
Uma imagem docker com base em postgres:11.2-alpine, para armazenar os dados de importação dos arquivos
### pdv-service
Serviço utilizando a plataforma spring-boot com jdk 8
Cria uma imagem docker com base em openjdk:8-jdk-alpine, ma qual compila e executa o pdv-service

Foi utlizado os seguintes comandos para geração dos certificados https:

```
  keytool -genkeypair -alias pdv -keyalg RSA -keysize 2048 -storetype PKCS12 -keystore pdv.p12 -validity 3650
```

```
  keytool -genkeypair -alias pdv -keyalg RSA -keysize 2048 -keystore pdv.jks -validity 3650
```

```
  keytool -importkeystore -srckeystore pdv.jks -destkeystore pdv.p12 -deststoretype pkcs12
```

## Endpoints
_**Atenção:**_ utilize o parametro -k na utilização do comando curl7
### uploadFile
Importa os arquivos

Exemplo de uso:
```
curl -X POST \
  https://localhost/uploadFile \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data' \
  -F data=@pdvs.csv
  -F tenant=<some-tenant-id>
``` 

### list
Metodo GET
Retorna a lista de pdvs com limit de 1000 registros

Parameters:
Name | Type | Mandatory | Description
------------ | ------------ | ------------ | ------------
tenantId | STRING | YES | identificador do tenant
nome | STRING | NO | nome do pdv
cidade | STRING | NO | cidade do pdv
endereco | STRING | NO | endereco do pdv
cep | STRING | NO | CEP do pdv
offset | INT | NO | offset da consulta, default: 0

Exemplo de uso:

```
  curl -k -X GET https://localhost:8443/list?tenantId=123&offset=2000
```
