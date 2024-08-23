## command
* api修改 `goctl api go -api app.api -dir .`
* 文档生成 `goctl api plugin -plugin goctl-swagger="swagger -filename doc/app.json" -api app.api -dir .`
* 文档预览 `docker run -d --name swag -p 8087:8080 -e SWAGGER_JSON=/opt/app.json -v /Users/dou/go/src/github.com/go-template/app/doc/:/opt swaggerapi/swagger-ui`
`DEV_MODE=app-prd dc up -d --force-recreate --build study-go`

## 环境变量
可以直接覆盖文件变量
