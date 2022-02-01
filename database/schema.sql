CREATE TABLE oauth_token
(
    service        TEXT                        NOT NULL,
    desmos_address TEXT                        NOT NULL,
    access_token   TEXT                        NOT NULL,
    refresh_token  TEXT                        NOT NULL,
    creation_time  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT single_oauth_token UNIQUE (service, desmos_address)
)