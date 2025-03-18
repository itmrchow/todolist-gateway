# todolist-gateway

# 簡介（Introduction）

- 定義與管理API
- 身份驗證與授權

## API 詳細規格（API Specification）

### API Resource

- `/users`: users 管理
- `/tasks`: tasks 管理

### Users Resource
| Method | Path            | Description |
| ------ | --------------- | ----------- |
| POST   | /users/register | 註冊新用戶  |
| POST   | /users/login    | 用戶登入    |


#### [POST] /users/register - 註冊新用戶
- request  
  - `email`: string , required , An email address
  - `name`: string , required , min 1 characters , max 20 characters
  - `password`: string , required , min 6 characters , max 20 characters , An alphanumeric string
``` json
{
  "email": "test@example.com",
  "name": "test",
  "password": "abc123"
}
```
- response (201 Created)
  - `message`: string , required 
    - `SUCCESS`: 註冊成功

``` json
{
  "message": "SUCCESS"
}
```

- response (400 Bad Request)
  - `message`: string , required 
    - `FAILED`: 格式不符合

``` json
{
  "message": "FAILED"
}
```

#### [POST] /users/login - 用戶登入
- request
  - `email`: string , required
  - `password`: string , required
- response (200 OK)
  - `message`: string , required
  - `token`: string , required
- response (400 Bad Request)
  - `message`: string , required
    - `FAILED`: 格式不符合
- response (401 Unauthorized)
  - `message`: string , required
    - `FAILED`: 登入失敗
- response (403 Forbidden)
  - `message`: string , required
    - `FAILED`: 找不到用戶

### Tasks Resource
| Method | Path              | Description  |
| ------ | ----------------- | ------------ |
| POST   | /tasks/create     | 新增任務     |
| PUT    | /tasks/update/:id | 更新任務     |
| DELETE | /tasks/delete/:id | 刪除任務     |
| GET    | /tasks/list       | 查詢任務列表 |


#### [POST] /tasks/create - 新增任務
- request  
  - `title`: string , required
  - `description`: string , required
  - `status`: string , required , enum: [pending, in_progress, completed]
- response (201 Created)
  - `message`: string , required
- response (400 Bad Request)

#### [PUT] /tasks/update/:id - 更新任務
- request
  - `title`: string , required
  - `description`: string , required
  - `status`: string , required , enum: [pending, in_progress, done]
- response (200 OK)
- response (400 Bad Request)

#### [DELETE] /tasks/delete/:id - 刪除任務  
- request
  - `id`: string , required
- response (200 OK)
- response (400 Bad Request)

#### [GET] /tasks/list - 查詢任務列表
- request
- response (200 OK)
- response (400 Bad Request)

### Common API response format
- 500 Internal Server Error
  - `message`: string , required
    - `SYSTEM_ERROR`: 系統錯誤
- 401 Unauthorized
  - `message`: string , required
    - `TOKEN_EXPIRED`: token 過期
    - `TOKEN_INVALID`: token 無效
- 403 Forbidden
  - `message`: string , required
    - `USAGE_LIMIT_EXCEEDED`: 使用限制超過
    - `UNAUTHORIZED_ACCESS`: 未授權
- 404 Not Found
  - `message`: string , required
    - `RESOURCE_NOT_FOUND`: 找不到資源
- 400 Bad Request
  - `message`: string , required
    - `INVALID_REQUEST`: 請求格式錯誤
