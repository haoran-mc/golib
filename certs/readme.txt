用于本地开发测试的证书与私钥


certs/
├── ec_cert.crt         # 自签证书，本地 HTTPS 测试
├── ec_private.key      # 私钥
└── ec_req.conf         # OpenSSL 配置


• 证书使用 ECC prime256v1 曲线
• 有效期 10 年
• 包含 subjectAltName：localhost, 127.0.0.1, ::1
