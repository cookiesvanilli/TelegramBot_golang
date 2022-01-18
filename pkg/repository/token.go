package repository

//разделение доступа по типам токена
type Bucket string

const (
	AccessTokens  Bucket = "access_tokens"
	RequestTokens Bucket = "request_tokens"
)

//request token и access token имеют один и тот же формат
type TokenRepository interface {
	//сохранить токен
	Save(chatID int64, token string, bucket Bucket) error
	//получть токена
	Get(chatID int64, bucket Bucket) (string, error)
}
