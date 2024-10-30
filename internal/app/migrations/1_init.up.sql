CREATE TABLE users (
                       id  serial NOT NULL PRIMARY KEY,
                       username text  NOT NULL UNIQUE,
                       password text  NOT NULL,
                       admin boolean not null DEFAULT false,
                       created_at 		timestamp with time zone 	DEFAULT now() NOT NULL,
                       updated_at 		timestamp with time zone
);

CREATE TABLE categories (
                            id  serial NOT NULL PRIMARY KEY,
                            name text  NOT NULL,
                            created_at 		timestamp with time zone 	DEFAULT now() NOT NULL,
                            updated_at 		timestamp with time zone
);

CREATE TABLE books (
                       id 	  serial NOT NULL PRIMARY KEY,
                       title text  NOT NULL,
                       author text  NOT NULL,
                       year integer not null CHECK (year > 0),
   stock integer not null CHECK (year >= 0),
   price integer not null CHECK (price > 0),
   category_id integer,
   created_at 		timestamp with time zone 	DEFAULT now() NOT NULL,
   updated_at 		timestamp with time zone,

   FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE carts (
                       user_id  integer NOT NULL primary key,
                       book_ids integer[] not null,
                       created_at 		timestamp with time zone 	DEFAULT now() NOT NULL,
                       updated_at 		timestamp with time zone,

                       FOREIGN KEY (user_id) REFERENCES users(id)
);