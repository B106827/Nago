--
-- Koki ASAUMI <roomikiru@gmail.com>
--

CREATE TABLE IF NOT EXISTS user (
  id               INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  email            VARCHAR(255) NOT NULL                COMMENT 'メールアドレス',
  password         VARCHAR(255) NOT NULL                COMMENT 'パスワード',
  line_user_id     VARCHAR(255)                         COMMENT 'LINE ユーザー ID',
  tel              VARCHAR(13)                          COMMENT '電話番号',
  name             VARCHAR(255)                         COMMENT '名前',
  address          VARCHAR(255)                         COMMENT '住所',
  `status`         INT NOT NULL DEFAULT 1               COMMENT 'ユーザーステータス',
  created_on       DATETIME NOT NULL                    COMMENT '登録日時',
  modified_on      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT 'ユーザー';

CREATE TABLE IF NOT EXISTS target (
  id               INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  user_id          INT NOT NULL                         COMMENT '管理者 ID',
  email            VARCHAR(255)                         COMMENT 'メールアドレス',
  tel              VARCHAR(13)                          COMMENT '電話番号',
  device_serial_no VARCHAR(100) NOT NULL                COMMENT 'Raspi シリアル番号',
  latest_passed_on DATETIME                             COMMENT '最終通過日時',
  name             VARCHAR(255) NOT NULL                COMMENT '名前',
  address          VARCHAR(255)                         COMMENT '住所',
  created_on       DATETIME NOT NULL                    COMMENT '登録日時',
  modified_on      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id),
  KEY target_key01 (user_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT 'ターゲット';

CREATE TABLE IF NOT EXISTS target_notification_conf (
  id                     INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  target_id              INT NOT NULL                         COMMENT 'ターゲット ID',
  weekday_id             INT NOT NULL                         COMMENT '曜日 ID',
  interval_min           INT NOT NULL                         COMMENT '間隔(分)',
  flg_every_notification INT NOT NULL DEFAULT 0               COMMENT '毎通知フラグ',
  flg_not_notification   INT NOT NULL DEFAULT 0               COMMENT '非通知フラグ',
  created_on             DATETIME NOT NULL                    COMMENT '登録日時',
  modified_on            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id),
  KEY target_notification_conf_key01 (target_id),
  KEY target_notification_conf_key02 (weekday_id, flg_every_notification)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT 'ターゲット・通知設定';

CREATE TABLE IF NOT EXISTS user_temporary (
  id            CHAR(10) NOT NULL COMMENT 'ID',
  email         VARCHAR(255)      COMMENT 'メールアドレス',
  oauth_user_id VARCHAR(255)      COMMENT 'OAuth ユーザー ID',
  user_id       INT UNSIGNED NULL COMMENT 'ユーザー ID (メールアドレス変更用)',
  redirect_path VARCHAR(255)      COMMENT 'リダイレクト先',
  created_on    DATETIME NOT NULL COMMENT '登録日時',
  modified_on   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最終更新日時',
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT 'ユーザー仮登録情報';

