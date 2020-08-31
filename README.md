# Test application
### Goals
Create a REST service where user can:
- register an account via login, password. log in; 
- fill out, change your profile;
- view the profile of other registered users;

##### Bonus:

- logging errors;
- the user can set himself an avatar. if the image is larger than 160x160px, the service should compress it;
- tests or code that is covered by tests with minimal modification (interfaces).
- at the discretion of the candidate: DB (RDBMS), ORM, images storage, router, libs

### Description technologies
- Database : `SQLite`
- Image Storage : `Minio`
- Auth : `JWT`

### What Done
- implement API for user
- login, signup
- resize and user's avatar

### Issues
- Docker doesn't run image with sqlite on `OSX`
- Error handling by type of errors
- What parameters should be config
- Error response without a message, only  http code
- Not ability to download avatar by link
- Pageable request works incorrect

### API's

#### 1. signup
##### request
```shell script
wget --no-check-certificate --quiet --method POST --timeout=0 --header 'Content-Type: application/json' \
  --body-data '{
  "email":"ytest@test.com",
  "password":"123123"
}
' \
   'http://127.0.0.1:8080/api/v1/user/signup'
``` 
##### response 
```json
{
    "userId": "c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc",
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTg4ODc0NjUsInVzZXJfaWQiOiJjMDIyZTljMy1jZjFmLTRkNzItODhiNi0xM2ZkZTRjZmI4YmMifQ.N0aJBGyGb74LM27zxf4LJMeuS1lsfnG27C-E1HbFAgo"
}
```

#### 2. login
##### request
```shell script
wget --no-check-certificate --quiet \
  --method POST \
  --timeout=0 \
  --header 'Content-Type: application/json' \
  --body-data '{
  "email":"ytest@test.com",
  "password":"123123"
}
' \
   'http://127.0.0.1:8080/api/v1/user/login'
```
##### response 
```json
{
    "user": {
        "userId": "c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc"
    },
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTg4ODc2NDksInVzZXJfaWQiOiJjMDIyZTljMy1jZjFmLTRkNzItODhiNi0xM2ZkZTRjZmI4YmMifQ.kinKyRCBGuZvqVtGKDpHYgvVd_UNf-TbdCsT5Qhz0ho"
}
```

#### 3. update profile
##### request 
```shell script
wget --no-check-certificate --quiet \
  --method PATCH \
  --timeout=0 \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTg4ODc0NjUsInVzZXJfaWQiOiJjMDIyZTljMy1jZjFmLTRkNzItODhiNi0xM2ZkZTRjZmI4YmMifQ.N0aJBGyGb74LM27zxf4LJMeuS1lsfnG27C-E1HbFAgo' \
  --header 'Content-Type: application/json' \
  --body-data '{
    "firstName": "Elon",
    "lastName":"Mask"
}' \
   'http://127.0.0.1:8080/api/v1/user/c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc/update'
```
##### response
```json
    {
        "userId": "c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc",
        "firstName": "Elon",
        "lastName": "Mask"
    }
```
#### 4. update avatar
##### request 
```shell script
# wget doesn't support file upload via form data, use curl -F \
wget --no-check-certificate --quiet \
  --method POST \
  --timeout=0 \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTg4ODc0NjUsInVzZXJfaWQiOiJjMDIyZTljMy1jZjFmLTRkNzItODhiNi0xM2ZkZTRjZmI4YmMifQ.N0aJBGyGb74LM27zxf4LJMeuS1lsfnG27C-E1HbFAgo' \
  --body-data '' \
   '127.0.0.1:8080/api/v1/user/c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc/update/avatar'
```

##### response 
```json
{
    "userId": "c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc",
    "firstName": "Elon",
    "lastName": "Mask",
    "avatarLink": "127.0.0.1:9000/user/avatar/c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc.jpeg"
}
```
#### 5. get user profile
##### request
```
wget --no-check-certificate --quiet \
  --method GET \
  --timeout=0 \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTg4ODg0MDAsInVzZXJfaWQiOiJjMDIyZTljMy1jZjFmLTRkNzItODhiNi0xM2ZkZTRjZmI4YmMifQ.JscnjNeJCL3pgc7WiDNKZYVElL9AHjDySy3H6jHyG9I' \
   'http://127.0.0.1:8080/api/v1/user/c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc'   
```

##### response 
```json
{
    "userId": "c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc",
    "firstName": "Elon",
    "lastName": "Mask",
    "avatarLink": "127.0.0.1:9000/user/avatar/c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc.jpeg"
}
```
#### 6. get another users
##### request 
```shell script
wget --no-check-certificate --quiet \
  --method GET \
  --timeout=0 \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTg4ODg4MDQsInVzZXJfaWQiOiJiYjM5MmZlNC0xNmIxLTQ0MTktYjdmNC05M2RmNDFmOWNkZjcifQ.yxXJ_iBGp_afw6H1fQS4xc9kkYmTh8wQEhHUCHKRQrQ' \
   'http://127.0.0.1:8080/api/v1/user/signed?pageSize=10&pageNumber=1'
```
##### response
```json
[
  {
    "userId": "c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc",
    "firstName": "Elon",
    "lastName": "Mask",
    "avatarLink": "127.0.0.1:9000/user/avatar/c022e9c3-cf1f-4d72-88b6-13fde4cfb8bc.jpeg"
  },
  {
    "userId": "5cbf96e8-d6f7-40cc-ba94-99712009004e",
    "firstName": "Elon2",
    "lastName": "Mask2"
  }
]
```
