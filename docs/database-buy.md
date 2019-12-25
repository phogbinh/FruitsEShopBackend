# Database discount policies tables
## cURL API
### make the transation
#### Description

make transation from cart

#### Authorization
not Required.
#### Response
200 when successful make transation

#### Example
`curl -v GET http://104.199.190.234:80/buy`

#### Get order
Get all order from `userName`
#### Authorization
not Required.
#### Response
A JSON object containing all product information
#### Example
curl -v http://localhost:8080/getorder?username=jamfly

#### Expected response format
```json
{
items: [
  {
    "DateTime":"2019-12-24 14:44:00",
     "Pname":"testProduct",
     "Price":100
  }]
}
```
