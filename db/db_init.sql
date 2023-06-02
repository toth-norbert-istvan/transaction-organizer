CREATE TABLE transaction_type (
  id            SERIAL PRIMARY KEY,
  name          VARCHAR(255) NOT NULL,
  type_group    VARCHAR(255) NOT NULL
);

CREATE TABLE transactions (
  id                       SERIAL PRIMARY KEY,
  partner                  VARCHAR(255) NOT NULL,
  amount                   NUMERIC(15, 2) NOT NULL,
  date                     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  transaction_type_id      INT,

  CONSTRAINT fk_transaction_type FOREIGN KEY(transaction_type_id) REFERENCES transaction_type(id)
);