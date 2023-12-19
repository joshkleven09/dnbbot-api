-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS hello
(
    id   UUID PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_profile
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id TEXT      NOT NULL UNIQUE,
    name        TEXT      NOT NULL,
    timezone    TEXT,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    deleted_at  TIMESTAMP
);

CREATE TABLE IF NOT EXISTS guild
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id TEXT      NOT NULL UNIQUE,
    name        TEXT      NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    deleted_at  TIMESTAMP
);

CREATE TABLE IF NOT EXISTS play_session
(
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_profile_id UUID      NOT NULL,
    guild_id        UUID      NOT NULL,
    start_time      TIMESTAMP NOT NULL,
    end_time        TIMESTAMP NOT NULL,
    game            TEXT,
    is_player       BOOLEAN   NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    deleted_at      TIMESTAMP,
    FOREIGN KEY (user_profile_id) REFERENCES user_profile (id),
    FOREIGN KEY (guild_id) REFERENCES guild (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS user_profile;
DROP TABLE IF EXISTS guild;
DROP TABLE IF EXISTS play_session;
DROP TABLE IF EXISTS hello;