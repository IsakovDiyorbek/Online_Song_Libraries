CREATE TABLE songs (
    id UUID PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,   -- Название группы
    song_name VARCHAR(255) NOT NULL,    -- Название песни
    release_date TIMESTAMP,                  -- Дата релиза песни
    text TEXT,                          -- Текст песни
    link VARCHAR(255),                  -- Ссылка на песню
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Дата создания записи
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Дата обновления записи
);
