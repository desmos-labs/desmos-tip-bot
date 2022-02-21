CREATE TABLE user_account
(
    id             SERIAL NOT NULL PRIMARY KEY,
    desmos_address TEXT   NOT NULL,
    CONSTRAINT unique_user_account UNIQUE (desmos_address)
);

CREATE TABLE user_preferences
(
    user_id  BIGINT  NOT NULL REFERENCES user_account (id) PRIMARY KEY,
    currency CHAR(3) NOT NULL
);

CREATE TABLE service_account
(
    id            SERIAL                      NOT NULL PRIMARY KEY,
    user_id       BIGINT                      NOT NULL REFERENCES user_account (id),
    service       TEXT                        NOT NULL,
    access_token  TEXT                        NOT NULL,
    refresh_token TEXT                        NOT NULL,
    creation_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT unique_service_account UNIQUE (service, user_id)
);

CREATE TABLE application_account
(
    service_account_id BIGINT NOT NULL REFERENCES service_account (id) ON DELETE CASCADE,
    application        TEXT   NOT NULL,
    username           TEXT   NOT NULL,
    CONSTRAINT unique_application_account UNIQUE (application, username)
);