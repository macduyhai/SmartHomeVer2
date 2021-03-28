package config

type Config struct {
	MySQLURL  string `envconfig:"MYSQL_URL" default:"root:1@tcp(localhost:3306)/qc_monitor?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"`
	APIKey    string `envconfig:"API_KEY" default:"duyhai-monitor-v1"`
	SecretKet string `envconfig:"SECRET_KEY" default:"duyhai-monitor-v1-key"`
}

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCF3e0oBwnmlxbb2ibwUhOanRgq
8M0zKBuMaz+lQokW8GrKHFbe6DDiPR9N93q2KW4da0T3hL03l4+BTOCRA2lZybDZ
xAMWgiQSdF85lH024T67J2y/BW/JHvD7PHbkM8OJ7NUtBPKHN57d+x7tLL6BenJV
a+hs5Lo7nE5IwbIurwIDAQAB
-----END PUBLIC KEY-----
`)

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCF3e0oBwnmlxbb2ibwUhOanRgq8M0zKBuMaz+lQokW8GrKHFbe
6DDiPR9N93q2KW4da0T3hL03l4+BTOCRA2lZybDZxAMWgiQSdF85lH024T67J2y/
BW/JHvD7PHbkM8OJ7NUtBPKHN57d+x7tLL6BenJVa+hs5Lo7nE5IwbIurwIDAQAB
AoGALBkYZ8gxzcFV6WSq0R3okVVQwcyDfGeo84/c7n7FlEXsl9nQECwi9lQ2PMPa
q6loOc69cGBMyMRnpKuDiqG6ETzxtKw3LPafGsXYD5FIZ34jOxRS2Ccsnwj84REj
Ww91iIli3kNZdqETV0Y0KXWkF7wMlQ9U597ZRmxa26BWxYECQQC8ASLpKK8ypRkc
C6c9/njZe0OpvLUkppQwe1YIfeMfm4EgzYy5BbjdAavsn9ANfE80A7OIxWqhkRZ7
ZkYhoFZBAkEAtkhUcntWM7ZPYvD033Q9JwAhEGf/Ciw0r+X5dwkwGPOoee7Mxd/z
zoBmgm0L2/57vrP204tF2vCjktht4Keo7wJAA9pGG98QkAogFJoMiFGxqktDXLQY
RjL/sGqmna/uupQWNlTgAF6kpirFmijAO7aDbP5ybGgXQk5V1puG7mN5wQJBAJJn
2dvxkDUMswqG+kcXt55Bjkz9Gm1zQAYfspSXPphr69+zm6k6zToJC0yqhSH3bjCn
nxIeBMdrDBZ/2xDb2OUCQHnot9aJVhXah5hTPKBG6fOrFLz3tvQiIDoqPmW2TH+g
62tZyiZylQ3xhWuPgF0pUbwSJJ46PZ52P6RCnDSMcTk=
-----END RSA PRIVATE KEY-----
`)
