# shopJsonRPC

### For start app use command ```docker-compose up -d```

## Methods:
  ### 1. GetAmount
   #### Request:
    curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "RpcApi.GetAmount", "params": [{"store":1}]}' http://localhost:8088/test
        
   #### Answer:
    {"id":1,"result":[
    {"storage_id":1,"product_id":1,"amount":7,"reserved":0},{"storage_id":1,"product_id":2,"amount":3,"reserved":0},
    {"storage_id":1,"product_id":3,"amount":9,"reserved":0},{"storage_id":1,"product_id":11,"amount":5,"reserved":0},
    {"storage_id":1,"product_id":12,"amount":6,"reserved":0},{"storage_id":1,"product_id":13,"amount":2,"reserved":0}],
    "error":null}

  ### 2. Reserve
   #### Request:
    curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "RpcApi.Reserve", "params": [{"ids":[1, 2]}]}' http://localhost:8088/test
   
   #### Answer:
    {"id":1,"result":"Success","error":null}
    
  ### 3. ReserveRelease
   #### Request:
    curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "RpcApi.ReserveRelease", "params": [{"ids":[1, 2]}]}' http://localhost:8088/test
   
   #### Answer:
    {"id":1,"result":"Success","error":null}    
