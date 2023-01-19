SET CHARSET UTF8;
SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;

/* user */
INSERT INTO user VALUES (1, 'roomikiru@gmail.com', '3b6c7aeaac7d2971c00e5e7401a61ee436d24df6', 1, NOW(), NOW());
/* product */
INSERT INTO product VALUES (1, '商品タイトル1商品タイトル1', '商品サブタイトル1商品サブタイトル1', 10000, '商品説明1商品説明1商品説明1商品説明1商品説明1商品説明1商品説明1商品説明1商品説明1商品説明1商品説明1商品説明1商品説明1', 1, NOW(), NOW());
INSERT INTO product VALUES (2, '商品タイトル2商品タイトル2', '商品サブタイトル2商品サブタイトル2', 10000, '商品説明2商品説明2商品説明2商品説明2商品説明2商品説明2商品説明2商品説明2商品説明2商品説明2商品説明2商品説明2商品説明2', 1, NOW(), NOW());
/* product_image */
INSERT INTO product_image VALUES (1, 1, 'https://d2air1d4eqhwg2.cloudfront.net/images/5959/500x500/c0d2e1c6-3e73-401c-9382-13cddb7106a4.jpg', NOW(), NOW());
INSERT INTO product_image VALUES (2, 1, 'https://d2air1d4eqhwg2.cloudfront.net/markdownx/9af7b408-179e-4e9b-b8d8-3aea9501c3a2.jpg', NOW(), NOW());
INSERT INTO product_image VALUES (3, 1, 'https://d2air1d4eqhwg2.cloudfront.net/markdownx/d726cb6e-46fd-4c25-8faf-0fe5679238bd.jpg', NOW(), NOW());
