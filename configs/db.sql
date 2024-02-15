

CREATE TABLE users (
  	id			varchar	NOT NULL,
  	name		VARCHAR(200)	NOT NULL,
	email		VARCHAR(100) NOT NULL,
	password	VARCHAR(100) NOT NULL,
	created_at		timestamp,
	updated_at		timestamp,
	PRIMARY KEY ( id )
);
CREATE UNIQUE INDEX u_user_email ON users(email);

CREATE TABLE user_address(
    id  varchar not null,
    user_id varchar not null,
    cep varchar not null,
    street varchar,
    city varchar,
    state varchar,
    PRIMARY KEY(id)
);
CREATE INDEX user_address_user_id ON user_address (user_id);
ALTER TABLE user_address ADD CONSTRAINT user_address_user_id FOREIGN KEY ( user_id ) REFERENCES users(id);

CREATE table event_types(
    id  varchar not null,
    description varchar not null,
    created_at		timestamp,
	updated_at		timestamp,
	PRIMARY KEY ( id )
); 

CREATE TABLE events(
  	id				varchar NOT NULL,
  	user_id			varchar NOT NULL,
	event_type_id   varchar not null,
	performed_at	timestamp,
	PRIMARY KEY ( id )
);	
CREATE INDEX events_user_id ON events (user_id);
ALTER TABLE events ADD CONSTRAINT events_user_id FOREIGN KEY ( user_id ) REFERENCES users(id);
CREATE INDEX events_event_type_id ON events (event_type_id);
ALTER TABLE events ADD CONSTRAINT events_event_type_id FOREIGN KEY ( event_type_id ) REFERENCES event_types(id);


				 