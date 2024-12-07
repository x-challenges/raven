module github.com/x-challenges/raven

go 1.23.4

godebug gotypesalias=1

replace github.com/imdario/mergo => dario.cat/mergo v1.0.0

require (
	github.com/99designs/gqlgen v0.17.57
	github.com/BurntSushi/toml v1.4.0
	github.com/TheZeroSlave/zapsentry v1.23.0
	github.com/aws/aws-sdk-go-v2 v1.32.6
	github.com/aws/aws-sdk-go-v2/config v1.28.6
	github.com/aws/aws-sdk-go-v2/service/sqs v1.37.2
	github.com/eko/gocache/lib/v4 v4.1.6
	github.com/eko/gocache/store/go_cache/v4 v4.2.2
	github.com/eko/gocache/store/redis/v4 v4.2.2
	github.com/getsentry/sentry-go v0.30.0
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.23.0
	github.com/go-redis/redis_rate/v10 v10.0.1
	github.com/go-resty/resty/v2 v2.16.2
	github.com/google/uuid v1.6.0
	github.com/graph-gophers/dataloader v5.0.0+incompatible
	github.com/hibiken/asynq v0.25.0
	github.com/json-iterator/go v1.1.12
	github.com/labstack/echo-contrib v0.17.1
	github.com/labstack/echo/v4 v4.12.0
	github.com/landrade/gqlgen-cache-control-plugin v1.1.0
	github.com/mcuadros/go-defaults v1.2.0
	github.com/nicksnyder/go-i18n/v2 v2.4.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/qmuntal/stateless v1.7.1
	github.com/redis/go-redis/v9 v9.7.0
	github.com/shopspring/decimal v1.4.0
	github.com/spf13/viper v1.19.0
	github.com/vektah/gqlparser/v2 v2.5.20
	github.com/ydb-platform/gorm-driver v0.2.0
	github.com/ydb-platform/ydb-go-sdk-auth-environ v0.5.0
	github.com/ydb-platform/ydb-go-sdk/v3 v3.93.1
	go.uber.org/fx v1.23.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f
	golang.org/x/text v0.20.0
	gorm.io/gorm v1.25.12
)

require (
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.47 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.21 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.2 // indirect
	github.com/aws/smithy-go v1.22.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gabriel-vasile/mimetype v1.4.7 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.5.1 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.60.1 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/sagikazarmark/locafero v0.6.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/yandex-cloud/go-genproto v0.0.0-20241202160409-e71211db0fd5 // indirect
	github.com/ydb-platform/ydb-go-genproto v0.0.0-20241112172322-ea1f63298f77 // indirect
	github.com/ydb-platform/ydb-go-yc v0.12.3 // indirect
	github.com/ydb-platform/ydb-go-yc-metadata v0.6.1 // indirect
	go.uber.org/dig v1.18.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/grpc v1.68.0 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/agnivade/levenshtein v1.2.0 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/gorilla/websocket v1.5.3
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/valyala/fasthttp v1.57.0
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.29.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/time v0.8.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)
