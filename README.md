# transaction-organizer

Aim of this project is Go self-learning.

The implemented application is a REST API that facilitates the organization of baking transactions and generates reports from collected data.

## Project setup
1. Database initialization
   - Install PostgreSQL server
   - Create `transaction-organizer` database
   - Execute `db/db_init.sql` SQL script to create DB schema
2. Start application with `main.go`
3. Postman collection is available here: [link](https://www.postman.com/payload-explorer-83303472/workspace/transaction-organizer/overview) 

## Upload transactions
Transactions can be imported with an Excel file. The currently supported Excel file format is exported by K&H Bank. You can upload it through the `{{baseUrl}}/transactions/kh?file-type=excel` POST endpoint.

All the included transactions are saved without transaction type, which can be set later via `...` PATCH endpoint.

## Get all transactions
Transactions can be fetched via `{{baseUrl}}/transactions` GET endpoint.