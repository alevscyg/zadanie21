CREATE TABLE IF NOT EXISTS todo (
  id bigserial not null primary key,
  userid bigserial,
  title varchar not null,
  description varchar
);