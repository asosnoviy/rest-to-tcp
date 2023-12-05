# Simple http2tcp proxy

Send TCP request over http

default port: 8050

### Request:  
HTTP request /simpletcp with json body
```json
{
    "Server":"192.168.40.29:1050", // tcp server adress and port
    "data":"Ak0DDQo=" // base64 sending data
}
```

### Answer:

```json
{
    "Error": "",
    "Data": "Ak1BVCAgICAgICxMICAwLjAsVyAgMC4wLEggIDAuMCxNLEtfX19fX18sRCAgMC4wMCxNLEY1MDAwLEkDDQo=",
}
```

## Example:
```cmd
curl --location --request GET '192.168.41.202:8050/simpletcp' \
--header 'Content-Type: application/json' \
--data '{
    "Server":"192.168.40.29:1050",
    "data":"Ak0DDQo="
}'
```

# HTTP сервис для перенаправления запросов в tcp

## Пример использования:

Запустить  
`rest-to-tcp.exe`  
или установить как службу  
`sc create rest2tcp binpath="C:\Program Files\rest-to-tcp\rest-to-tcp.exe" DisplayName="Rest to TCP proxy"`


```bsl

	Данные = "{
		| "server":"192.168.40.29:1050", // Адрес и порт tcp устройства
		| "data":"Ak0DDQo=" // base64 данные для отправки
		|}";

	Соединение = Новый HTTPСоединение("192.168.41.202", 8050); // Адрес запущенного tcpproxy
	
	Заголовки = Новый Соответствие();
	Заголовки.Вставить("Content-Type", "application/json");
	Запрос = Новый HTTPЗапрос("/simpletcp", Заголовки);
	
	Запрос.УстановитьТелоИзСтроки(Данные);

	Ответ = Соединение.ОтправитьДляОбработки(Запрос);
	ТелоОтвета = Ответ.ПолучитьТелоКакСтроку();
	
	ЧтениеJSON = Новый ЧтениеJSON;
	ЧтениеJSON.УстановитьСтроку(ТелоОтвета);
	ДанныеОтвет = ПрочитатьJSON(ЧтениеJSON);

	ДвоичныеДанные = ПолучитьДвоичныеДанныеИзBase64Строки(ДанныеОтвет.data); // Ответ от устройства
	
```
