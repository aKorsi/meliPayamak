<div dir="rtl">

# سورس اتصال به سامانه پیامکی [ملی پیامک](https://www.melipayamak.com/) 

روش استفاده:

<div dir="ltr">

```go

mpService := meliPayamak.RestService{
			UserName: username,
			PassWord: password,
		}

resp,err := mpService.SendSMS(to, from, text, isFlash)
or
resp,err := mpService.SendByBaseNumber(text, to, bodyId)

```

</div>

در صورت خطا یا مشکلی حتما اعلام بفرمایید 

</div>
