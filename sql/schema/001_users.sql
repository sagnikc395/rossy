-- goose runs the migrations in order 
-- they have a up and a down statement 
-- up -> create, down -> delete
-- helps to easily rollup and down the changes

-- +goose up
CREATE TABLE users (
    id UUID primary key, 
    created_at timestamp not null,
    updated_at timestamp not null ,
    name TEXT not null 
);

-- +goose down 
DROP table users;