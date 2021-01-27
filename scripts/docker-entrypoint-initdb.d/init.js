db = db.getSiblingDB('fl-auth')
db.createUser(
{
  user: "fl-auth",
  pwd: "fl-auth123", 
  roles: [
    { role: "readWrite", db: "fl-auth" }
  ]
})
