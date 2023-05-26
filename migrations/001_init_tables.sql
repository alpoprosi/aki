-- +goose Up

create table users
(
    id                    serial primary key,
    uuid                  character varying(36),

    name        text not null, 
    email       text not null,
    phone       varchar(25) not null,
    password    text not null,
    last_name   text default null,
    patronic    text default null,
    is_landLord boolean,
    is_admin    boolean,

    created_at            timestamp with time zone default now() not null,
    updated_at            timestamp with time zone default now() not null,
    deleted_at            timestamp with time zone default null,
    
    constraint users_uuid unique (uuid)
);

create index idx_users_uuid on users using btree (uuid);

create table categories
(
    id                    serial primary key,
    name                  text not null,
    created_at            timestamp with time zone default now() not null,
    updated_at            timestamp with time zone default now() not null,
    deleted_at            timestamp with time zone default null,

    constraint categorie_name unique (name)
);

create index idx_categories_names on categories (name);

create table landlords
(
    id                    serial primary key,
    category_id           serial reference categories on (id),

    juridical_name  text not null,
    registrar_job   text not null,
    inn             text not null,
    descriprion     text not null,
    lable_reference text not null,

    created_at            timestamp with time zone default now() not null,
    updated_at            timestamp with time zone default now() not null,
    deleted_at            timestamp with time zone default null,
);

create index idx_landlords_categories on landlords (category_id);
create index idx_landlords_inn on landlords (inn);
create index idx_landlords_juridical on landlords (juridical_name);


-- +goose Down

drop table users, landlords, categories;
