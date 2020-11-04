# Dockerized-Golang-Postgres-Mysql-API

Follow this guide about how it is done:
https://levelup.gitconnected.com/dockerized-crud-restful-api-with-go-gorm-jwt-postgresql-mysql-and-testing-61d731430bd8


docker build -t full_app .

source .env


curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDQ0NDkxNTcsInVzZXJfaWQiOjF9.pijf9Dy91mULJLe2nxA0pht0QkYEzN0fzWnILt8Ttqg" \
     -H "Content-Type: application/json" \
     -X POST \
     --data '{"title": "testtitle", "content": "yolo", "author_id": 1}' \
     localhost:8080/api/v1/post
{"id":3,"title":"testtitle","content":"yolo","author":{"id":1,"nickname":"loeken","email":"loeken@internetz.me","password":"$2a$10$zKhxvG4SyVIXDhZuOdHWl.kY4zf9ASfLUrG/FHsTCTQ5qc32rQrj2","created_at":"2020-11-03T23:18:36.78785Z","updated_at":"2020-11-03T23:18:36.78785Z"},"author_id":1,"created_at":"2020-11-03T23:20:07.543236915Z","updated_at":"2020-11-03T23:20:07.543237145Z"}
