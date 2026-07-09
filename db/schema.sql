--MODULO 1


CREATE TABLE usuarios (
    id          INTEGER PRIMARY KEY,
    nombre      TEXT    NOT NULL,
    email       TEXT    NOT NULL UNIQUE,
    password    TEXT    NOT NULL,
    nivel       INTEGER NOT NULL,
    reputacion  INTEGER NOT NULL
);

CREATE TABLE reviews (
    id            INTEGER PRIMARY KEY,
    reviewer_id   INTEGER NOT NULL,
    reviewed_id   INTEGER NOT NULL,
    rating        INTEGER NOT NULL,
    comment       TEXT    NOT NULL,
    FOREIGN KEY (reviewer_id) REFERENCES usuarios(id),
    FOREIGN KEY (reviewed_id) REFERENCES usuarios(id)
);

CREATE TABLE badges (
    id            INTEGER PRIMARY KEY,
    nombre        TEXT    NOT NULL,
    descripcion   TEXT    NOT NULL,
    required_rep  INTEGER NOT NULL
);


--MODULO 2


CREATE TABLE categorias (
    id      INTEGER PRIMARY KEY,
    nombre  TEXT    NOT NULL
);

CREATE TABLE productos (
    id             INTEGER PRIMARY KEY,
    nombre         TEXT    NOT NULL,
    descripcion    TEXT    NOT NULL,
    precio         REAL    NOT NULL,
    categoria_id   INTEGER NOT NULL,
    FOREIGN KEY (categoria_id) REFERENCES categorias(id)
);

CREATE TABLE ordenes (
    id             INTEGER PRIMARY KEY,
    producto_id    INTEGER NOT NULL,
    comprador_id   INTEGER NOT NULL,
    estado         TEXT    NOT NULL,
    FOREIGN KEY (producto_id) REFERENCES productos(id),
    FOREIGN KEY (comprador_id) REFERENCES usuarios(id)
);


--MODULO 3

CREATE TABLE mensajes (
    id            INTEGER PRIMARY KEY,
    sender_id     INTEGER NOT NULL,
    receiver_id   INTEGER NOT NULL,
    content       TEXT    NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES usuarios(id),
    FOREIGN KEY (receiver_id) REFERENCES usuarios(id)
);

CREATE TABLE misiones (
    id               INTEGER PRIMARY KEY,
    title            TEXT    NOT NULL,
    description      TEXT    NOT NULL,
    required_level   INTEGER NOT NULL,
    reward_points    INTEGER NOT NULL
);

CREATE TABLE user_missions (
    id            INTEGER PRIMARY KEY,
    user_id       INTEGER NOT NULL,
    mission_id    INTEGER NOT NULL,
    completed     BOOLEAN NOT NULL,
    FOREIGN KEY (user_id) REFERENCES usuarios(id),
    FOREIGN KEY (mission_id) REFERENCES misiones(id)
);