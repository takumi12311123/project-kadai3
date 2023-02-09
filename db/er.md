```mermaid
erDiagram

  users {
  id string 
  name string 
  email string 
  password string 
  created_at datetime
  updated_at datetime
  organization_id organization
  }
  
  organizations {
   id string
   name string
   created_at string
   updated_at string
   user_id users
  }
  
  users }|..|{ organizations: company
```
