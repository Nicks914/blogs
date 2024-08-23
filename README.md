# blogs
A simple blog platform where users can create accounts, log in, and create, read, update, and delete blog posts.


## Curl Commands

1. Register a New User

curl -X POST http://localhost:8080/api/register \
-H "Content-Type: application/json" \
-d '{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "password": "password123"
}'

2. Login

curl -X POST http://localhost:8080/api/login \
-H "Content-Type: application/json" \
-d '{
  "email": "john.doe@example.com",
  "password": "password123"
}'

3. View User Profile (Authenticated)

curl -X GET http://localhost:8080/api/profile \
-H "Authorization: Bearer <JWT_TOKEN>"

4. Update User Profile (Authenticated)

curl -X PUT http://localhost:8080/api/profile/update \
-H "Authorization: Bearer <JWT_TOKEN>" \
-H "Content-Type: application/json" \
-d '{
  "name": "John Updated"
}'

5. Create a New Blog Post (Authenticated)

curl -X POST http://localhost:8080/api/posts/create \
-H "Authorization: Bearer <JWT_TOKEN>" \
-H "Content-Type: application/json" \
-d '{
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post."
}'

6. Get All Blog Posts

curl -X GET http://localhost:8080/api/posts

7. Get a Single Blog Post by ID

curl -X GET http://localhost:8080/api/posts?id=1

8. Update a Blog Post (Authenticated)

curl -X PUT http://localhost:8080/api/posts/update?id=1 \
-H "Authorization: Bearer <JWT_TOKEN>" \
-H "Content-Type: application/json" \
-d '{
  "title": "Updated Blog Post",
  "content": "This is the updated content."
}'

9. Delete a Blog Post (Authenticated)

curl -X DELETE http://localhost:8080/api/posts/delete?id=1 \
-H "Authorization: Bearer <JWT_TOKEN>"
