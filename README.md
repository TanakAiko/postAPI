# postAPI
A go API for the handling of posts within a forum

## Before starting
import package: 
- [sqlite3](github.com/mattn/go-sqlite3) : `go get github.com/mattn/go-sqlite3`

## Note
The stucture of the body of the request is the following :
- ### created post
```
{
    action: "createPost"
    body: {
	    userID       int   
        categorie    []string
	    content      string
    }
}
```

- ### get one post
```
{
    action: "getOne"
    body: {
	    id      int
    }
}
```

- ### get all post
```
{
    action: "getAll"
}
```

- ### delete post
```
{
    action: "delete"
    body: {
	    id      int
    }
}
```

## Testing
### createPost
- #### request
Execute the following command :
```
curl -X POST http://localhost:8082/ -d '{
  "action":"createPost", 
  "body": 
    { 
      "userID": 1,
      "categorie": ["Manga", "Anime", "Berserk"],
      "content": "I am th black swordman"
    }
}' -H "Content-Type: application/json"
```
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusCreated (201)
- body    : "New post created"
```

### getOne
- #### request
Execute the following command :
```
curl -X POST http://localhost:8082/ -d '{
  "action":"getOne", 
  "body": 
    { 
      "id": 1
    }
}' -H "Content-Type: application/json"
```
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusOK (200)
- body    : the post data
```

### getAll
- #### request
Execute the following command :
```
curl -X POST http://localhost:8082/ -d '{
  "action":"getAll"
}' -H "Content-Type: application/json"
```
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusOK (200)
- body    : all the post data
```

### delete
- #### request
Execute the following command :
```
curl -X POST http://localhost:8082/ -d '{
  "action":"delete", 
  "body": 
    { 
      "id": 1
    }
}' -H "Content-Type: application/json"
```
- #### response
If the resquet is well executed, the response should be :
```
- status  : http.StatusOK (200)
- body    : "Post well deleted"
