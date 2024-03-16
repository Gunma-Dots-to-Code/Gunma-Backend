# Gunma-Backend

## 動作確認
1. `.env.example`→`.env`にリネームします。また特有の環境変数があれば埋めます
2. dockerでgoの実行環境を用意します
```bash
docker compose up -d
```
3. 実行環境に入ります
```bash
docker exec -it gunma-backend /bin/bash
```
4. デバッグします
```bash
go run main.go
```

## ER図
```mermaid
erDiagram
    COMPANY {
        string company_id PK
        string domain
    }
    ADMIN {
        string admin_id PK
        string email
        string password
        string company_id FK
    }
    RESPONDER {
        string responder_id PK
        string email
        string password
        string company_id FK
    }
    USER {
        string user_id PK
        string email
        string password
        string username
    }
    QUESTION {
        string question_id PK
        datetime date
        string content
        string user_id FK
        string category_id FK
    }
    ANSWER {
        string answer_id PK
        string question_id FK
        datetime date
        string content
        string user_id FK
    }
    CATEGORY {
        string category_id PK
        string company_id FK
        string name
    }
    CATEGORY ||--o{ CATEGORY : "parent"
    COMPANY ||--|{ ADMIN : "has"
    COMPANY ||--o{ RESPONDER : "has"
    COMPANY ||--|{ CATEGORY : "has"
    USER ||--o{ QUESTION : "creates"
    USER ||--o{ ANSWER : "provides"
    QUESTION ||--o{ ANSWER : "is answered by"
    CATEGORY ||--o{ QUESTION : "categorizes"
```
