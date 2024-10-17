version: "0.1"
database:
  # consult[https://gorm.io/docs/connecting_to_the_database.html]"
  dsn: "host=localhost user=debill password=debill dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  # mysql or postgres
  db  : "postgres"
  # enter the required data table or leave it blank.
  tables: ["account", "address", "customer", "location"]
  # specify a directory for output
  outPath :  "../pkg/gen/model"
  # query code file name, default: gen.go
  outFile :  "model.go"
  # generate unit test for query code
  withUnitTest  : false
  # generated model code's package name
  modelPkgName  : "model"
  # generate with pointer when field is nullable
  fieldNullable : false
  # generate field with gorm index tag
  fieldWithIndexTag : false
  # generate field with gorm column type tag
  fieldWithTypeTag  : false
