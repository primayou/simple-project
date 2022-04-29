CREATE TABLE transactions (
    id serial not null primary key,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    currency TEXT NOT NULL,
    amount TEXT NOT NULL
);
