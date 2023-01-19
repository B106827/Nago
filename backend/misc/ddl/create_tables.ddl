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
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT 'ユーザー';

CREATE TABLE IF NOT EXISTS user_temporary (
  id         CHAR(10) NOT NULL                                                        COMMENT 'ID',
  email      VARCHAR(255)                                                             COMMENT 'メールアドレス',
  expired_at DATETIME NOT NULL                                                        COMMENT '有効期限日時',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT 'ユーザー仮登録情報';

CREATE TABLE IF NOT EXISTS product (
  id          INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  name        VARCHAR(255) NOT NULL                                                    COMMENT '商品名',
  sub_title   VARCHAR(255)                                                             COMMENT 'サブタイトル',
  price       INT UNSIGNED NOT NULL                                                    COMMENT '価格',
  description TEXT                                                                     COMMENT '商品説明',
  status      INT NOT NULL DEFAULT 1                                                   COMMENT '商品ステータス',
  created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT '商品';

CREATE TABLE IF NOT EXISTS product_image (
  id         INT UNSIGNED NOT NULL AUTO_INCREMENT                                     COMMENT 'ID',
  product_id INT UNSIGNED NOT NULL                                                    COMMENT '商品ID',
  url        TEXT                                                                     COMMENT '画像URL',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT '作成日時',
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT '商品画像';

