# newbilling-client-go

Go client for [qingcloud newbilling system](https://beatflow.qingcloud.com/)

To generate go client:
- download and install go-swagger cli from https://github.com/go-swagger/go-swagger/releases
- download https://api.beatflow.qingcloud.com/swagger-ui/api.swagger.json to ./openapi/swagger.json
- edit ./openapi/swagger.json and add `"x-nullable": true` to all date-time (`"format": "date-time"`) fields.
- run `rm -rf client models && swagger generate client -f ./openapi/swagger.json --skip-validation && go mod tidy`
