ERD E-WALLET

```mermaid
erDiagram
direction LR
    user ||--o{ balance : own
    user ||--o{ pin : own
    user ||--o{ topup : do
 user ||--o{ transaction : do
    transaction ||--o{ history : have

    user {
        int id_user PK
        string name
        string email
        string img
        string password
    }
    pin {
        int id_pin PK
        int pin
        int id_user FK
    }
    balance {
        int id_balance PK
        int id_user FK
        int saldo
        date last_updated
    }


    transaction {
        int id_transaction PK
        int amount

        enum status "pending,completed,failed"
        date date_transaction
        string method
        int id_user FK
    }
    topup {
        int id_topup PK
        int amount
        date date_transaction
        int id_user FK
    }

    history {
        int id_history PK
        int id_transaction FK
    }


```
