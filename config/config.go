package config

import "github.com/spf13/viper"

type AppConfig struct {
	Env              string           `mapstructure:"environment"`
	Port             string           `mapstructure:"port"`
	ServiceName      ServiceName      `mapstructure:"service-name"`
	Secrets          Secrets          `mapstructure:"secrets"`
	BinanceFutureUrl BinanceFutureUrl `mapstructure:"binance-future-url"`
}

type Secrets struct {
	BinanceApiKey    string `mapstructure:"binance-apiKey"`
	BinanceSecretKey string `mapstructure:"binance-secretKey"`
}

type ServiceName struct {
	BinanceFuture string `mapstructure:"binance-future"`
	BinanceSpot   string `mapstructure:"binance-spot"`
}

type BinanceFutureBaseUrl struct {
	BianceUrl1 string `mapstructure:"binance1"`
	// BianceUrl2 string `mapstructure:"binance2"`
	// BianceUrl3 string `mapstructure:"binance3"`
	// BianceUrl4 string `mapstructure:"binance4"`
}

type BinanceFutureUrl struct {
	SetLeverage          string               `mapstructure:"set-leverage"`
	SingleOrder          string               `mapstructure:"single-order"`
	MultipleOrder        string               `mapstructure:"miltiple-order"`
	BinanceFutureBaseUrl BinanceFutureBaseUrl `mapstructure:"binance-future-baseUrl"`
}

func ReadConfig() (c *AppConfig, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.Unmarshal(&c)

	if c.Env == "local" {
		c.BinanceFutureUrl.BinanceFutureBaseUrl.BianceUrl1 = "https://testnet.binancefuture.com"
		// c.Binance.SpotBaseUrl = "https://testnet.binance.vision"
	}

	return c, nil
}
