CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY,
    payer_id UUID NOT NULL,
    reciver_id UUID NOT NULL,
    value NUMERIC(10, 2) NOT NULL,
    FOREIGN KEY (payer_id) REFERENCES users(id),
    FOREIGN KEY (reciver_id) REFERENCES users(id)
);