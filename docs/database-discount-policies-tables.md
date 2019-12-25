# Database discount policies tables
## cURL API
### Create staff discount policy
#### Description
Create a discount policy for the given staff to the database table `discount_policies` and to either the database table `shipping_discount_policies`, `seasonings_discount_policies` or `special_event_discount_policies` according to the given discount policy type. Additionally, the database table `products` will also be updated should the given discount policy is of type Special Event.
#### Authorization
Required.
#### Response
A JSON object containing the newly created discount policy's information fetched from the database.
#### Example
##### Create discount policy `SHI000001` by staff `bill`
`curl -X POST -H "Authorization: Bearer <token>" -d "{ \"code\": \"SHI000001\", \"name\": \"A Shipping Discount Policy\", \"description\": \"Description\", \"type\": \"Shipping\", \"staffUserName\": \"bill\", \"shippingMinimumOrderPrice\": \"500.00\", \"seasoningsRate\": \"\", \"seasoningsBeginDate\": \"\", \"seasoningsEndDate\": \"\", \"specialEventRate\": \"\", \"specialEventBeginDate\": \"\", \"specialEventEndDate\": \"\", \"specialEventProductIds\": [] }" localhost:8080/auth/bill/discount-policies`
##### Create discount policy `SEA000001` by staff `bill`
`curl -X POST -H "Authorization: Bearer <token>" -d "{ \"code\": \"SEA000001\", \"name\": \"A Seasonings Discount Policy\", \"description\": \"Description\", \"type\": \"Seasonings\", \"staffUserName\": \"bill\", \"shippingMinimumOrderPrice\": \"\", \"seasoningsRate\": \"0.05\", \"seasoningsBeginDate\": \"2019-01-01\", \"seasoningsEndDate\": \"2020-01-01\", \"specialEventRate\": \"\", \"specialEventBeginDate\": \"\", \"specialEventEndDate\": \"\", \"specialEventProductIds\": [] }" localhost:8080/auth/bill/discount-policies`
##### Create discount policy `SPE000001` by staff `bill` (products having IDs `10` and `25` must be in the database table `products` before calling the API)
`curl -X POST -H "Authorization: Bearer <token>" -d "{ \"code\": \"SPE000001\", \"name\": \"A Special Event Discount Policy\", \"description\": \"Description\", \"type\": \"Special Event\", \"staffUserName\": \"bill\", \"shippingMinimumOrderPrice\": \"\", \"seasoningsRate\": \"\", \"seasoningsBeginDate\": \"\", \"seasoningsEndDate\": \"\", \"specialEventRate\": \"0.05\", \"specialEventBeginDate\": \"2018-01-01\", \"specialEventEndDate\": \"2021-01-31\", \"specialEventProductIds\": [\"10\", \"25\"] }" localhost:8080/auth/bill/discount-policies`
#### Expected response format
```json
{
    "code": "SHI000001",
    "name": "A Shipping Discount Policy",
    "description": "Description",
    "type": "Shipping",
    "staffUserName": "bill",
    "shippingMinimumOrderPrice": "500.00",
    "seasoningsRate": "",
    "seasoningsBeginDate": "",
    "seasoningsEndDate": "",
    "specialEventRate": "",
    "specialEventBeginDate": "",
    "specialEventEndDate": ""
}
```
#### Error response format
```json
{
    "error": "error message."
}
```

### Get staff discount policies
#### Description
Get all discount policies' information of the given staff from the database.
#### Authorization
Required.
#### Response
A JSON object containing all discount policies' information of the given staff from the database.
#### Example
`curl -X GET -H "Authorization: Bearer <token>" localhost:8080/auth/bill/discount-policies`
#### Expected response format
```json
[
    {
        "code": "SPE000001",
        "name": "A Special Event Discount Policy",
        "description": "Description",
        "type": "Special Event",
        "staffUserName": "bill",
        "shippingMinimumOrderPrice": "",
        "seasoningsRate": "",
        "seasoningsBeginDate": "",
        "seasoningsEndDate": "",
        "specialEventRate": "0.05",
        "specialEventBeginDate": "2018-01-01",
        "specialEventEndDate": "2021-01-31",
    },
    {
        "code": "SEA000001",
        "name": "A Seasonings Discount Policy",
        "description": "Description",
        "type": "Seasonings",
        "staffUserName": "bill",
        "shippingMinimumOrderPrice": "",
        "seasoningsRate": "0.05",
        "seasoningsBeginDate": "2019-01-01",
        "seasoningsEndDate": "2020-01-01",
        "specialEventRate": "",
        "specialEventBeginDate": "",
        "specialEventEndDate": "",
    },
    {
        "code": "SHI000001",
        "name": "A Shipping Discount Policy",
        "description": "Description",
        "type": "Shipping",
        "staffUserName": "bill",
        "shippingMinimumOrderPrice": "500.00",
        "seasoningsRate": "",
        "seasoningsBeginDate": "",
        "seasoningsEndDate": "",
        "specialEventRate": "",
        "specialEventBeginDate": "",
        "specialEventEndDate": ""
    }
]
```
#### Error response format
```json
{
    "error": "error message."
}
```

### Delete staff discount policy
#### Description
Delete a discount policy of the given staff from the database table `discount_policies` and cascadingly from either the database table `shipping_discount_policies`, `seasonings_discount_policies` or `special_event_discount_policies` according to the request-to-be-deleted discount policy's type. Additionally, should the given discount policy's type be Special Event, associated products will also be updated to set null on their corresponding fields.
#### Authorization
Required.
#### Response
N/A.
#### Example
`curl -X DELETE -H "Authorization: Bearer <token>" localhost:8080/auth/bill/discount-policies/SPE000001`
#### Error response format
```json
{
    "error": "error message."
}
```