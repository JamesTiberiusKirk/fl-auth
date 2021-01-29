db = db.getSiblingDB('users')
db.createUser(
{
  user: "fl-auth",
  pwd: "fl-auth123",
  roles: [
    { role: "readWrite", db: "users" }
  ]
})