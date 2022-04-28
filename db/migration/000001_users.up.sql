CREATE TABLE "user" (
    id serial not null primary key,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL
);
