Приложение после запуска начнет генерировать сообщения и отправлять их в Redis. 
Каждое сообщение имеет указанные в genAt - время создания, в genBy - id генератора и в text - непосредственный текст сообщения.
Каждое запущенное приложение обзаводится уникальным случайным id. Текст сообщения генерируется случайно. 
Для отправки в Redis сообщение формируется в JSON.

Пример генерируемого сообщения:

    Message has been generated:  {"GenAt":"11-26-2018 07:25:49 Mon","GenBy":"xVtWyHvz","Text":"Voluptatem accusantium sit perferendis aut consequatur."}

Запустите еще одно или несколько приложений и они начнут проверять сообщения, помещенные в Redis. 

		The following message has been SUCCESSFULLY verified: {"GenAt":"11-26-2018 07:28:20 Mon","GenBy":"xVtWyHvz","Text":"Accusantium voluptatem consequatur aut perferendis sit."}

С заданной вероятностью 5% проверяемые сообщения будут определяться как сообщения с ошибками. 
Такие сообщения отправляются в Redis в список "errors". 

    The following message is INCORRECT:  {"GenAt":"11-26-2018 07:26:23 Mon","GenBy":"xVtWyHvz","Text":"Aut voluptatem accusantium perferendis consequatur sit."}
		

