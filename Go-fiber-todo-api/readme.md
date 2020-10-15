## ENDPOINTS

### Signup - POST /api/users/signup

Signup functionality, requires a body with email and password

```
{
	"email":"verte.fra@gmail.com",
	"password":"verte"
}
```

**response object**

```
{
    "ok": true,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InZlcnRlLmZyYUBnbWFpbC5jb20iLCJleHAiOjE2MDI3OTUwNDh9.yyKhnuI-95rMWIkb_f1mG8O9iPs8oaqjmO082SDAJWw",
    "userID": {
        "_id": "5f888aa3972ec4acade4e707"
    }
}
```

- Encrypted in the token you can find the email of the user. Store the Id in the state or in local storage to make requests for the todo related to that user

**Error Messages**

- If the email is already in use

```
{
    "email": "verte.fra@gmail.com",
    "error": "email already in use",
    "ok": false
}
```

- email or password field empty

```
{
    "error": "empty fields",
    "ok": false
}
```

### Login - POST /api/users/login

- Login functionality, requires a body with email or password

```
{
	"email":"verte.fra@gmail.com",
	"password":"verte"
}
```

**response object**

```
{
    "ok": true,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InZlcnRlLmZyYUBnbWFpbC5jb20iLCJleHAiOjE2MDI3OTUwNDh9.yyKhnuI-95rMWIkb_f1mG8O9iPs8oaqjmO082SDAJWw",
    "userID": {
        "_id": "5f888aa3972ec4acade4e707"
    }
}
```

**Error Messages**

- Email or password field empty

```
{
    "error": "empty fields",
    "ok": false
}
```

- Email not found

```
{
    "error": "mongo: no documents in result",
    "ok": false
}
```

- Wrong password

```
{
    "error": "wrong password",
    "ok": false
}

```

### GetAllTodos - GET api/todos?user=userID

- Returns all the todo objects related to the user with **userID**

- Response object

```
{
    "ok": true,
    "todos": [
        {
            "_id": "5f88a946d5c8c43ef62765d1",
            "created_at": "2020-10-15T19:55:50.662Z",
            "updated_at": "2020-10-15T19:55:50.662Z",
            "userID": "5f888aa3972ec4acade4e707",
            "title": "first todo",
            "description": "todo",
            "done": false
        }
    ]
}
```

### GetTodoByID - GET api/todos/:id

- Returns a specific todo object with id equal to :id

- Response object

```
{
    "ok": true,
    "todo": {
        "_id": "5f88a946d5c8c43ef62765d1",
        "created_at": "2020-10-15T19:55:50.662Z",
        "updated_at": "2020-10-15T19:55:50.662Z",
        "userID": "5f888aa3972ec4acade4e707",
        "title": "first todo",
        "description": "todo",
        "done": false
    }
}
```
