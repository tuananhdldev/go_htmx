create table todos(
    id serial primary key,
    todo varchar(128),
    done boolean
);
insert into todos(todo,done) values('hello!',false);