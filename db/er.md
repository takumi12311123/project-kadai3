```mermaid
erDiagram

  users {
  string id
  string name
  string email
  string password
  datetime created_at
  datetime updated_at
  }
  
  organizations {
   string id
   string name
   datetime created_at
   datetime updated_at
  }
  
  users }|..|{ organizations
```
