package localize

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

// Service localize interface
type Service interface {
	// DefaultLangiage returns default language code
	DefaultLanguage() LanguageCode

	// Languages returns available language
	Languages() []LanguageCode

	// Localize returns translated string with provided options
	Localize(lang LanguageCode, messageID string, opts ...Option) (string, error)

	// MustLocalize returns translated string with provided options or raise panic
	MustLocalize(lang LanguageCode, messageID string, opts ...Option) string
}

// Service interface implementation
type service struct {
	logger     *zap.Logger
	options    *Config
	languages  []LanguageCode
	bundle     *i18n.Bundle
	localizers map[LanguageCode]*i18n.Localizer
}

// NewService constructor
func NewService(logger *zap.Logger, options *Config) Service {
	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// load message files
	for _, lang := range options.Langs {
		bundle.MustLoadMessageFile(
			filepath.Join(
				options.Path,
				fmt.Sprintf("active.%s.toml", strings.ToLower(lang)),
			),
		)
	}

	var (
		languages  = MustParseLanguageCodes(options.Langs...)
		localizers = make(map[LanguageCode]*i18n.Localizer)
	)

	// register localizers
	for _, code := range languages {
		localizers[code] = i18n.NewLocalizer(
			bundle,
			[]string{
				code,
				DefaultLanguageCode,
			}...,
		)
	}

	return &service{
		logger:     logger,
		options:    options,
		languages:  languages,
		bundle:     bundle,
		localizers: localizers,
	}
}

// Languages return slice of available language codes
func (s *service) Languages() []LanguageCode {
	return s.languages
}

// DefaultLanguage implement Service interface
func (s *service) DefaultLanguage() LanguageCode {
	return DefaultLanguageCode
}

func (s *service) localizer(lang LanguageCode) *i18n.Localizer {
	return s.localizers[lang]
}

// Localize implement Service interface
func (s *service) Localize(lang LanguageCode, messageID string, opts ...Option) (string, error) {
	o := &Options{}

	for _, opt := range opts {
		opt(o)
	}

	output, err := s.localizer(lang).Localize(
		&i18n.LocalizeConfig{
			MessageID:    messageID,
			TemplateData: o.TemplateData,
			PluralCount:  o.PluralCount,
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "localization failed")
	}
	return output, nil
}

// MustLocalize implement Service interface
func (s *service) MustLocalize(lang LanguageCode, messageID string, opts ...Option) (output string) {
	output, err := s.Localize(lang, messageID, opts...)
	if err != nil {
		s.logger.Panic("must localization failed", []zap.Field{
			zap.String("language", lang),
			zap.String("message_id", messageID),
			zap.Any("options", opts),
			zap.Error(err),
		}...)
	}
	return
}
