## ENDPOINTS

### /api/users/signup

Signup functionality, requires a body with email and password

```{
	"email":"verte.fra@gmail.com",
	"password":"verte"
    }
```

**response object**

```{
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

```{
    "email": "verte.fra@gmail.com",
    "error": "email already in use",
    "ok": false
}
```

- email or password field empty

```{
    "error": "empty fields",
    "ok": false
}
```

### /api/users/login

Login functionality, requires a body with email or password

```{
	"email":"verte.fra@gmail.com",
	"password":"verte"
    }
```

**response object**

```{
    "ok": true,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InZlcnRlLmZyYUBnbWFpbC5jb20iLCJleHAiOjE2MDI3OTUwNDh9.yyKhnuI-95rMWIkb_f1mG8O9iPs8oaqjmO082SDAJWw",
    "userID": {
        "_id": "5f888aa3972ec4acade4e707"
    }
}
```

**Error Messages**

- Email or password field empty

```{
    "error": "empty fields",
    "ok": false
}
```

- Email not found

```{
    "error": "mongo: no documents in result",
    "ok": false
}
```

- Wrong password

```{
    "error": "wrong password",
    "ok": false
}

```
