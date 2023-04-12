--
-- Koki ASAUMI <roomikiru@gmail.com>
--

-- creatte_at,updated_at は ORMに合わせてカラム設定しているので注意

CREATE TABLE IF NOT EXISTS user (
  id         INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  email      VARCHAR(255) NOT NULL                                                    COMMENT 'メールアドレス',
  password   VARCHAR(255) NOT NULL                                                    COMMENT 'パスワード',
  name       VARCHAR(255) NOT NULL                                                    COMMENT '名前',
  status     INT NOT NULL DEFAULT 1                                                   COMMENT 'ユーザーステータス',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id),
  KEY user_index01 (email, password, status)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT 'ユーザー';

CREATE TABLE IF NOT EXISTS user_temporary (
  id         CHAR(10) NOT NULL                                                        COMMENT 'ID',
  email      VARCHAR(255)                                                             COMMENT 'メールアドレス',
  expired_at DATETIME NOT NULL                                                        COMMENT '有効期限日時',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT 'ユーザー仮登録情報';

CREATE TABLE IF NOT EXISTS product (
  id          INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  name        VARCHAR(255) NOT NULL                                                    COMMENT '商品名',
  sub_title   VARCHAR(255)                                                             COMMENT 'サブタイトル',
  price       INT UNSIGNED NOT NULL                                                    COMMENT '価格',
  description TEXT                                                                     COMMENT '商品説明',
  stock       INT UNSIGNED NOT NULL DEFAULT 0                                          COMMENT '在庫数',
  status      INT NOT NULL DEFAULT 1                                                   COMMENT '商品ステータス',
  created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '商品';

CREATE TABLE IF NOT EXISTS product_image (
  id         INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  product_id INT UNSIGNED NOT NULL                                                    COMMENT '商品ID',
  url        TEXT                                                                     COMMENT '画像URL',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id)
  KEY product_image_index01 (product_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '商品画像';

CREATE TABLE IF NOT EXISTS cart (
  id         INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  user_id    INT UNSIGNED NOT NULL                                                    COMMENT 'ユーザーID',
  product_id INT UNSIGNED NOT NULL                                                    COMMENT '商品ID',
  num        INT UNSIGNED NOT NULL                                                    COMMENT '数量',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id),
  UNIQUE KEY cart_unique01 (user_id, product_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT 'カート';

CREATE TABLE IF NOT EXISTS `order` (
  id          INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  user_id     INT UNSIGNED NOT NULL                                                    COMMENT 'ユーザーID',
  total_price INT UNSIGNED NOT NULL                                                    COMMENT '注文金額',
  status      INT UNSIGNED NOT NULL DEFAULT 0                                          COMMENT 'ステータス（0:未決済、1:決済済み, 2:決済失敗)',
  ordered_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP                              COMMENT '注文日時',
  created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id),
  KEY order_index01 (user_id, status_flg)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '注文情報';

CREATE TABLE IF NOT EXISTS order_detail (
  id         INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  order_id   INT UNSIGNED NOT NULL                                                    COMMENT '注文情報ID',
  product_id INT UNSIGNED NOT NULL                                                    COMMENT '商品ID',
  price      INT UNSIGNED NOT NULL                                                    COMMENT '商品価格（注文時点）',
  num        INT UNSIGNED NOT NULL                                                    COMMENT '数量',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id),
  KEY order_detail_index01 (order_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '注文情報詳細';

CREATE TABLE IF NOT EXISTS order_delivery_info (
  id                     INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  user_id                INT UNSIGNED NOT NULL                                                    COMMENT 'ユーザーID',
  order_id               INT UNSIGNED NOT NULL                                                    COMMENT '注文情報ID',
  last_name              VARCHAR(255) NOT NULL                                                    COMMENT '名字',
  first_name             VARCHAR(255) NOT NULL                                                    COMMENT '名前',
  primary_postcode       VARCHAR(3) NOT NULL                                                      COMMENT '郵便番号（郵便区番号）',
  secondary_postcode     VARCHAR(4) NOT NULL                                                      COMMENT '郵便番号（町域番号）',
  pref_id                INT UNSIGNED NOT NULL                                                    COMMENT '都道府県ID',
  primary_address        VARCHAR(255) NOT NULL                                                    COMMENT '住所1',
  secondary_address      VARCHAR(255)                                                             COMMENT '住所2',
  primary_phone_number   VARCHAR(4) NOT NULL                                                      COMMENT '電話番号1',
  secondary_phone_number VARCHAR(4) NOT NULL                                                      COMMENT '電話番号2',
  third_phone_number     VARCHAR(4) NOT NULL                                                      COMMENT '電話番号3',
  created_at             TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at             TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id),
  KEY order_delivery_info_index01 (user_id, order_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '注文情報 > 配送先情報';
