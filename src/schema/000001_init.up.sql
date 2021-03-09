CREATE TABLE cats
(
    id   serial       not null unique,
    age  int          not null,
    name varchar(255) not null
);

-- migrate -path ./schema -database 'postgres://postgres:123456@localhost:5432/catdb?sslmode=disable' up