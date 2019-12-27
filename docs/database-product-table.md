# Database `product` table
### Add product
#### Description
新增一個 product 到 `product` table，ProductId為流水號，從 1000 開始
#### Response
Status code
#### Input
```
StaffUserName
Description
Pname
Category
Source
Price
Inventory
SoldQuantity
OnSaleDate
ImageSrc

---
Post
http://localhost:8080/addproduct?StaffUserName=bill&Description=Description&Pname=apple&Category=fruit&Source=taiwan&Price=100&Inventory=5&SoldQuantity=10000&OnSaleDate=2019-12-25&ImageSrc=file

```
#### Example
```
/addproduct?StaffUserName=bill&Description=Description&Pname=apple...
```
#### Expected response format
```
Status code : 200
```
#### Error response format
```
Status code : 400、417
```

### Delete product
#### Description
從`product` table刪除一個product
#### Response
Status code
#### Input
```
ProductId

---
DELETE
http://localhost:8080/deleteproduct?ProductId=1000

```
#### Example
```
/deleteproduct?ProductId=1
```
#### Expected response format
```
Status code : 200
```
#### Error response format
```
Status code : 400、417
```

### query product
#### Description
從`product`取得指定的product
#### Response
Status code + json object
#### Input
```

Pname
或
StaffUserName

---
GET
http://localhost:8080/queryproduct?Pname=apple

```
#### Rule
能夠透過商品名稱搜尋到商品，也能利用用戶的名字來搜尋，**優先利用用戶名稱搜尋**，搜尋不到則回傳空json。

#### Example
```
/queryproduct?Pname=apple
```
#### Expected response format
```json
{
    "items":
    [
        {
            "ProductId":1000
            "PName":"Gold apple", 
            "SUser":"StaffA"
            "Category":"Apple", 
            "Description":"Very good apple!", 
            "Source":"Taiwan", 
            "Price":100, 
            "Inventory":20,
            "Quantity":1,
            "SaleDate":"2019-01-01"
        }
        ...
    ]
}
```
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

### Modify product
#### Description
修改 `product` table內的 productInfo
#### Response
Status code
#### Input
```
ProductId
StaffUserName
Description
Pname
Category
Source
Price
Inventory
SoldQuantity
OnSaleDate
ImageSrc

```
#### Example
```
/modifyproduct/ProductId=1&s_username=jeff&description=is good!&p_name=apple...
```
#### Expected response format
```
Status code : 200
```
#### Error response format
```
Status code : 400、403
```