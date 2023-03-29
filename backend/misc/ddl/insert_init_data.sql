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

/* pref_master */
INSERT INTO pref_master (name) VALUES ('北海道'), ('青森県'), ('岩手県'), ('宮城県'), ('秋田県'), ('山形県'), ('福島県'), ('茨城県'), ('栃木県'), ('群馬県'), ('埼玉県'), ('千葉県'), ('東京都'), ('神奈川県'), ('新潟県'), ('富山県'), ('石川県'), ('福井県'), ('山梨県'), ('長野県'), ('岐阜県'), ('静岡県'), ('愛知県'), ('三重県'), ('滋賀県'), ('京都府'), ('大阪府'), ('兵庫県'), ('奈良県'), ('和歌山県'), ('鳥取県'), ('島根県'), ('岡山県'), ('広島県'), ('山口県'), ('徳島県'), ('香川県'), ('愛媛県'), ('高知県'), ('福岡県'), ('佐賀県'), ('長崎県'), ('熊本県'), ('大分県'), ('宮崎県'), ('鹿児島県'), ('沖縄県');
