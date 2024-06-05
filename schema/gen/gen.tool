version: "0.1"
database:
  # consult[https://gorm.io/docs/connecting_to_the_database.html]"
  dsn : "postgres:postgres@tcp(localhost:5432)/debill?charset=utf8mb4&parseTime=true&loc=Local"
  # input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]
  db  : "postgres"
  # enter the required data table or leave it blank.
  tables  : "event,label,actor,user,provider,customer,account,address,catalog,product,offert,subscriptiob,subscription_history,fee,service,usage,billing,cycle,invoice,charge,ledger,tax,payment,cash"
  # specify a directory for output
  outPath :  "../../pkg/gen/model"
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
