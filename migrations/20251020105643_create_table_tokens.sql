-- +goose Up
create table tokens(
    id serial primary key,
    user_id integer not null,
    foreign key (user_id)  references users(id) on delete cascade,
    token varchar(2000) not null
);


-- +goose Down
DROP TABLE tokens;
