INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$aBvM/wZYS7A36L4YwcUWxunkxf1PHuKIoeeLU2ZtIUfjDuVUg/sLO"),
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$aBvM/wZYS7A36L4YwcUWxunkxf1PHuKIoeeLU2ZtIUfjDuVUg/sLO"),
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$aBvM/wZYS7A36L4YwcUWxunkxf1PHuKIoeeLU2ZtIUfjDuVUg/sLO");


INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES
(1, 2),
(1, 3),
(3, 1);

INSERT INTO publicacoes (titulo, conteudo, autor_id)
VALUES
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Oba!", 3);
