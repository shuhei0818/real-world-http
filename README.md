# Read World HTTP

[Real World HTTP](https://www.oreilly.co.jp/books/9784873119038/)の写経



## Hash Algorism

| Algorism                         | Explanation                                                                                                                                                          |
| -------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| MD5 ( Message Digest 5 )         | 128ビットのハッシュ関数。セキュリティ用途では非推奨                                                                                                                  |
| SHA-1（Secure Hash Algorithm 1） | 160ビットのハッシュ関数。セキュリティ用途では非推奨                                                                                                                  |
| SHA-2（Secure Hash Algorighm 2） | SHA-224, SHA-256, SHA-384, SHA-512, SHA-512/224, SHA-512/256の6つのバリエーションがある。<br> SHA-512が最も安全性が高く、一般的にはSHA-256が最もよく利用されている。 |
| SHA-3（Secure Hash Algorighm 3） | 最新のハッシュ関数。                                                                                                                                                 |

### e.g.
```bash
# MD5
$ md5 4-2-1-hash-algorism/sample.txt
> MD5 (4-2-1-hash-algorism/sample.txt) = 5177a58dc65c8a14dc90c69db3bf3dd2 # 16進表記の長さ: 32文字

# SHA-256
shasum -a 256 4-2-1-hash-algorism/sample.txt 
aaf9ff488e0767da5ea1d56118e6f65a16c5633b0cefc1fa089bd3ab1810613d  4-2-1-hash-algorism/sample.txt # 16進表記の長さ: 64文字
```
