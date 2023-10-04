CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL ,
    surname  VARCHAR(255) NOT NULL
);

CREATE TABLE accounts
(
    id             SERIAL PRIMARY KEY,
    currency_code  VARCHAR(3),
    active_balance FLOAT DEFAULT 0,
    frozen_balance FLOAT DEFAULT 0,
    user_id        INT,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE transactions
(
    id                   SERIAL PRIMARY KEY,
    currency_code        VARCHAR(3),
    transaction_status   VARCHAR(25),
    sender_account_id    INT,
    recipient_account_id INT,
    amount               FLOAT,
    transaction_date     TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (sender_account_id) REFERENCES accounts (id),
    FOREIGN KEY (recipient_account_id) REFERENCES accounts (id)
);
