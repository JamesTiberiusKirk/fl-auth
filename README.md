# FL-Auth Service

## Endpoints

### POST /register
- Payload
```json
{
    "email":"test@email.com",
    "username":"test",
    "password":"testpass"
}
```


### POST /login
- Payload
```json
{
    "email":"test@email.com",
    "password":"testpass"
}
```

### POST /verify_me
- Payload
```json
{
    "jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZW1haWwuY29tIiwicm9sZXMiOlsidXNlciJdLCJleHAiOjE2MTE3NDI4MTIsImlzcyI6ImZsLWF1dGgifQ.Xkv5QE6GPS4XyC6-QHXRn0E1chBq9qpOTNxxvRBObwk"
}
```