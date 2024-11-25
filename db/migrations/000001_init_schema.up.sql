CREATE TABLE  categories (
    category_id serial PRIMARY KEY,
    category_name character varying(15) NOT NULL,
    description text,
    picture bytea
);