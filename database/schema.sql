CREATE TABLE user
(
    id             SERIAL PRIMARY,
    desmos_address TEXT NOT NULL UNIQUE,
);

CREATE TABLE service_account
(
    user_id       BIGINT NOT NULL REFERENCES user (id),
    service       TEXT   NOT NULL,
    access_token  TEXT   NOT NULL,
    refresh_token TEXT   NOT NULL,
    creation_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT single_oauth_token UNIQUE (service, user_id)
);

CREATE TABLE application_account
(
    oauth_token_id BIGINT NOT NULL REFERENCES oauth_token (id),
    application    TEXT   NOT NULL,
    username       TEXT   NOT NULL,
    CONSTRAINT unique_username UNIQUE (application, username)
);