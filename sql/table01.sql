CREATE TABLE users (
    id      uuid,
    name    varchar(50),
    cpf     varchar(15),   
    email   varchar(35),   
    phone   varchar(20),
    PRIMARY KEY(id)
);

CREATE TABLE cards (
    id          uuid,
    id_user     uuid,
    brand       varchar(15),
    number      varchar(20),
    alias       varchar(50),
    exp_year    int,
    exp_month   int,
    PRIMARY KEY(id),
    CONSTRAINT fk_user
        FOREIGN KEY(id_user)
            REFERENCES users(id)
);




-- insert into client values(1, 'Alberto');
-- insert into client values(2, 'Fernando');
