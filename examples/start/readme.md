# Starting Example

Schema taken from: <https://github.com/ent/ent/tree/master/examples/start>

## Schema

> **Note**
>
> The following schema was generated by `entmaid`.

<!-- #start:entmaid -->
```mermaid
erDiagram
 Car {
  int id PK
  string model
  time-Time registered_at
  int user_cars FK
 }

 Group {
  int id PK
  string name
 }

 group_users {
  int group_id PK,FK
  int user_id PK,FK
 }

 User {
  int id PK
  int age
  string name
 }

 Group |o--o{ group_users : users-groups
 User |o--o{ Car : cars-owner
 User |o--o{ group_users : groups-users

```
<!-- #end:entmaid -->