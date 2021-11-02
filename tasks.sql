create table tasks(
    id serial primary key,
    name varchar(100),
    description varchar(255),
    created_at timestamp,
    updated_at timestamp
)