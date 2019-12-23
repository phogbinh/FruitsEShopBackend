# Database `order_item` table
### Add order item
#### Description
新增一個order item到`order_item` table
#### Response
Status code
#### Input
```
ProductId
CartId
Quantity
```
#### Example
```
/addorderitemtocart/ProductId=1&CartId=2&Quantity=3
```
#### Expected response format
```
Status code : 200
```
#### Error response format
```
Status code : 400、417
```

### Delete order item
#### Description
從`order_item` table刪除一個order item
#### Response
Status code
#### Input
```
ProductId
CartId
```
#### Example
```
/deleteorderitemincart/ProductId=1&CartId=2
```
#### Expected response format
```
Status code : 200
```
#### Error response format
```
Status code : 400、417
```

### Get all order items in cart
#### Description
從`order_item`取得目前在購物車內的所有order item
#### Response
Status code + json object
#### Input
```
N/A
```
#### Example
```
/getorderitemsincart/ProductId=1&CartId=2
```
#### Expected response format
```json
{
    "items":
    [
        {
            "PName":"Gold apple", 
            "Category":"Apple", 
            "Description":"Very good apple!", 
            "Source":"Taiwan", 
            "Price":100, 
            "Inventory":20
        }, 
        {
            "PName":"Pupu apple", 
            "Category":"Apple", 
            "Description":"Very bad apple!", 
            "Source":"Taiwan", 
            "Price":1, 
            "Inventory":200
        }
        ...
    ]
}```
#### Error response format
```json
{
    "code":400, 
    "items":
    []
}
```
or
```json
{
    "code":403, 
    "items":
    []
}
```

### Modify order item
#### Description
修改`order_item` table內的order item
#### Response
Status code
#### Input
```
ProductId
CartId
Quantity
```
#### Example
```
/modifyorderitemquantity/ProductId=1&CartId=2&Quantity=4
```
#### Expected response format
```
Status code : 200
```
#### Error response format
```
Status code : 400、403
```