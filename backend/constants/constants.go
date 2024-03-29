package constants

const (
	/*
	  サービス関連
	*/
	SERVICE_NAME = "Nago"

	/*
	  ユーザー関連
	*/
	// アクティブユーザーのステータス
	USER_STATUS_REGISTERED = 1

	/*
	  セッション関連
	*/
	// JWTのCookieキー名
	SESSION_COOKIE_KEY_NAME = "nago-cs"
	// CSRFのContextキー名
	CSRF_CONTEXT_KEY_NAME = "nago-ctc"
	// CSRFのCookieキー名
	CSRF_COOKIE_KEY_NAME = "nago-cc"

	/*
	  時間関連
	*/
	// 1日（秒）
	TIME_SECONDS_A_DAY = 86400

	/*
	  商品関連
	*/
	// 有効な商品のステータス
	PRODUCT_STATUS_ACTIVE = 1

	/*
	  Stripe・注文関連
	*/
	ORDER_CURRENCY     = "jpy"
	ORDER_SUCCESS_PATH = "/order/success"
	ORDER_CANCEL_PATH  = "/order/cancel"
)
