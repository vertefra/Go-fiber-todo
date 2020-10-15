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
- **Needs authentication token**

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
- **Needs authentication token**

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

### CreateTodo - POST api/todos?user=userID

- Creates a new todo Object associated with the user **userID**
- **Needs authentication token**
- Requires a body object with **title** and **description**

```
{
	"title": "new todo from today",
	"description":"brief description"
}
```

- Response object

```
{
    "ok": true,
    "todo": {
        "_id": "5f88b406665a6fdcb1e5d0e1",
        "created_at": "2020-10-15T20:41:42.902361913Z",
        "updated_at": "2020-10-15T20:41:42.902363876Z",
        "userID": "5f888aa3972ec4acade4e707",
        "title": "new todo from today",
        "description": "brief description",
        "done": false
    }
}
```

- TODO: verifiy the id in the query is associated with a user

**Error Messages**

- No user id in the url query

```
{
    "error": "no user id found",
    "ok": false
}
```

### UpdateTodo - PATCH /api/todos/:id

- Updates a todo obeject with id == :id
- **Needs authentication token**
- Needs a body object with the updated fields

```
{
    "title": "Changed",
    "description":"Changed",
    "done": true
}
```

- Response object

```
{
    "ok": true,
    "updated": {
        "_id": "5f88a946d5c8c43ef62765d1",
        "created_at": "2020-10-15T19:55:50.662Z",
        "updated_at": "2020-10-15T20:53:37.080688641Z",
        "userID": "5f888aa3972ec4acade4e707",
        "title": "Changed",
        "description": "Changed",
        "done": true
    }
}
```

**Error Messages**

- todo id is invalid

```
{
    "error": "encoding/hex: odd length hex string",
    "ok": false
}
```

### UpdateTodo - DELETE /api/todos/:id

- Deletes a todo obeject with id == :id
- **Needs authentication token**

- Response Object

```
{
    "deleted": {
        "_id": "5f88a946d5c8c43ef62765d1",
        "created_at": "2020-10-15T19:55:50.662Z",
        "updated_at": "2020-10-15T20:53:37.08Z",
        "userID": "5f888aa3972ec4acade4e707",
        "title": "Changed",
        "description": "Changed",
        "done": true
    },
    "ok": true
}
```

Next things todo:

- Converting ObjectID into string and send id in token
  token will contains the id and every operation will check that
  the id of the user and the userID for the todo is the same to
  restrict the data that an user can manipulate
