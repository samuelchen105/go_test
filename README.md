# go_test

## 第一部分

### 使用環境
* docker 20.10.6
* docker-compose 1.28.2
* docker image go:latest(1.16)
* docker image mysql:5.7

### 使用套件
* [gin-gonic/gin](https://github.com/gin-gonic/gin)
* [gorm.io/gorm](https://github.com/go-gorm/gorm)
* [gorm.io/driver/mysql](https://github.com/go-gorm/mysql)

### 完成的部分
* 使用Docker佈署
* 設計MYSQL資料表
* 實作管理商品目錄的CRUD RESTful API
* 實作管理商品的CRUD RESTful API

### 缺少的部分
* 資料驗證
* 實作進階的RESTful API
* testing

### 還原步驟
請先確保已安裝docker、docker-compose和git
```
git clone https://github.com/yuhsuan105/go_test.git
cd go_test/part1

# 若docker還沒啟動
sudo service docker start
sudo docker-compose up -d
##############################################################
由於我沒有設置讓app等待資料庫的腳本，第一次up可能會失敗，請再試一次
##############################################################
sudo docker-compose ps
# 若app和db的state都為up則成功建置完成

# 執行完後清理
sudo docker-compose down -rmi all
```

### API
```
# Create Directory
POST url:port/api/dir(ex: localhost:8080/api/dir)
body form: {name=?, is_hide=?}

# Read One Directory
GET url:port/api/dir/:id(ex: localhost:8080/api/dir/1)

# Read All Directories
GET url:port/api/dir/all(ex: localhost:8080/api/dir/all)

# Update Directory
PATCH url:port/api/dir(ex: localhost:8080/api/dir)
body form: {id=?, name=?, is_hide=?}

# Delete Directory
DELETE url:port/api/dir/:id(ex: localhost:8080/api/dir/1)

--------------------

# Create Product
POST url:port/api/pd(ex: localhost:8080/api/pd)
body form: {name=?, is_hide=?}

# Read One Product
GET url:port/api/pd/:id(ex: localhost:8080/api/pd/1)

# Read All Products
GET url:port/api/pd/all(ex: localhost:8080/api/pd/all)

# Update Product
PATCH url:port/api/pd(ex: localhost:8080/api/pd)
body form: {id=?, name=?, is_hide=?}

# Delete Product
DELETE url:port/api/pd/:id(ex: localhost:8080/api/pd/1)
```
img資料夾內有我使用postman測試時的截圖

## 第二部分

### 使用環境
* go 1.13.8

### 還原步驟
只有main.go檔

```
cd part2
go run main.go
```