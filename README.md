# 百度语音合成

安装：
```
git clone https://github.com/rocket049/baiduTTS.git
cd baiduTTS
go build .
```

**运行程序时需要`baiduTTS`目录中的`tmpls`，因此请在`baiduTTS`目录中运行本程序。**


**使用本软件需要一个百度大脑语音应用ID。**

首先把百度大脑语音应用ID写入下面形式的JSON文件（例如：app.json）：
```
{
	"AppID":"xxxxxx",
	"ApiKey":"xxxxxxxxxxxxxx",
	"SecretKey":"xxxxxxxxxxxxxxxxxxxxxx"
}
```

然后用下面的参数运行程序：
```
./baiduTTS -i /path/to/app.json
```

**运行后会自动打开浏览器，在页面输入框中输入文字，点击“合成”即可，注意文字数不可过多，字数最好不超过2000个，否则可能失败。**
