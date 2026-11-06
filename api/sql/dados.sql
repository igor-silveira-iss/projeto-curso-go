insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$aBvM/wZYS7A36L4YwcUWxunkxf1PHuKIoeeLU2ZtIUfjDuVUg/sLO"),
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$aBvM/wZYS7A36L4YwcUWxunkxf1PHuKIoeeLU2ZtIUfjDuVUg/sLO"),
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$aBvM/wZYS7A36L4YwcUWxunkxf1PHuKIoeeLU2ZtIUfjDuVUg/sLO");


insert into seguidores (usuario_id, seguidor_id)
values
(1, 2),
(1, 3),
(3, 1);
