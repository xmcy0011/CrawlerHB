# API

## 获取验证码
```json
{"storagePosition": "database"}
```

```json
{
    "xLenght": 227,
    "yLenght": 80,
    "picNum": "1",
    "timeTemple": "1581848594250",
    "storagePosition": null
}
```

## 校验验证码
```json
{"timeTemple": "1580804094410", "xLenght": 218, "storagePosition": "database"}
```

```text
true
```

## 登录
```json
{"suLoginid": "public", "suPasswd": "New@098Hero%", "sdId": "d211c9993e4844cebe510db6d4a7bd7d"}
```

```json
{
    "success": true,
    "flag": "Data",
    "msg": "登录成功",
    "data": {
        "url": "psInfo/psInfo.html",
        "result": null,
        "name": null,
        "regionCode": null,
        "sessUser": "4a2c1f4091444292ad7e242c2cc3ef97",
        "userLevel": null
    }
}
```

## 废气排放
```json
{"porttypeCode":2,"psid":"39bd36aac0c946b7a55033ec5e3d5f2d"}
```
 
 ```json
{
    "success": true,
    "flag": "Data",
    "msg": "成功!",
    "data": true
}
```

## 获取工厂所有监测点列表
```json
{"portTypeCode": 2, "psId": "85a15c8477fa452eb50876c4f4b84964", "startTime": "2019-04-11 00", "endTime": "2019-04-11 14"}
```

```json
{
    "success": true,
    "flag": "ListData",
    "msg": "查询成功",
    "data": [
        {
            "color": "#000000",
            "psinfoid": "130400000245",
            "name": "1号2号焦炉脱硫后排口",
            "id": "9",
            "dgi_mn": "1304270L0CX029",
            "monitorpointtype_code": "2"
        },
        {
            "color": "#000000",
            "psinfoid": "130400000245",
            "name": "1号焦炉脱硫前排口",
            "id": "1",
            "dgi_mn": "130427YTJH0001",
            "monitorpointtype_code": "5"
        },
        {
            "color": "#000000",
            "psinfoid": "130400000245",
            "name": "2号焦炉脱硫前排口",
            "id": "2",
            "dgi_mn": "1304270L0CX028",
            "monitorpointtype_code": "5"
        },
        {
            "color": "#000000",
            "psinfoid": "130400000245",
            "name": "地面站",
            "id": "11",
            "dgi_mn": "130427YTJH0002",
            "monitorpointtype_code": "5"
        }
    ]
}
```

## 获取某个监测点历史数据的表头信息
```json
{"colStr":"0", "periodType":"2061", "pollutionType":2, "psId":"85a15c8477fa452eb50876c4f4b84964", "outletNo":"11","choose":"1"}
```

```json
{
    "success": true,
    "flag": "ListData",
    "msg": "成功!",
    "data": [
        {
            "id": "date",
            "name": "监测时间",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "01-avg",
            "name": "烟尘实测浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "01-zsavg",
            "name": "烟尘折算浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "02-avg",
            "name": "二氧化硫实测浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "02-zsavg",
            "name": "二氧化硫折算浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "03-avg",
            "name": "氮氧化物实测浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "03-zsavg",
            "name": "氮氧化物折算浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "S01-avg",
            "name": "氧含量均值(%)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "S02-avg",
            "name": "烟气流速平均值(m/s)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": false,
            "ordering": false,
            "children": null
        },
        {
            "id": "S03-avg",
            "name": "烟气温度平均值(℃)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": false,
            "ordering": false,
            "children": null
        },
        {
            "id": "S04-avg",
            "name": "烟气动压平均值(兆帕)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": false,
            "ordering": false,
            "children": null
        },
        {
            "id": "S05-avg",
            "name": "烟气湿度平均值(%)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": false,
            "ordering": false,
            "children": null
        },
        {
            "id": "S06-avg",
            "name": "制冷温度平均值(℃)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": false,
            "ordering": false,
            "children": null
        },
        {
            "id": "S07-avg",
            "name": "烟道截面积(㎡)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": false,
            "ordering": false,
            "children": null
        },
        {
            "id": "S08-avg",
            "name": "烟气压力平均值(兆帕)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": false,
            "ordering": false,
            "children": null
        },
        {
            "id": "is_stop",
            "name": "是否停运",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        }
    ]
}
```

## 获取某个监测点历史数据
```json
{
	"length":5,
	"search":{
		"psId":"85a15c8477fa452eb50876c4f4b84964",
		"pollutionType":"2",
		"outletNo":"9",
		"colStr":"0",
		"periodType":"2061",
		"fromTime":"2020-02-01 00",
		"toTime":"2020-02-01 23",
		"header":"",
		"choose":"1",
		"data_source":"1"
	},
	"start":"1"
}
```

```json
{
    "success": true,
    "flag": "PageData",
    "msg": "查询成功！",
    "data": [
        {
            "date": "2020-02-01 00:00:00",
            "03-zsavg": "62.30$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "62.30$50.48291666666666",
            "02-zsavg": "1.08$3.3470833333333334",
            "S03-avg": "63.11$58.86958333333331",
            "01-zsavg": "2.33$1.6291666666666667",
            "S01-avg": "13.35$11.805000000000001",
            "S05-avg": "3.12$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "2.33$1.6291666666666667",
            "02-avg": "1.08$3.3470833333333334",
            "S08-avg": "-87.87$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.10$6.945416666666667"
        },
        {
            "date": "2020-02-01 01:00:00",
            "03-zsavg": "61.04$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "61.04$50.48291666666666",
            "02-zsavg": "1.58$3.3470833333333334",
            "S03-avg": "62.61$58.86958333333331",
            "01-zsavg": "2.46$1.6291666666666667",
            "S01-avg": "13.25$11.805000000000001",
            "S05-avg": "3.30$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "2.46$1.6291666666666667",
            "02-avg": "1.58$3.3470833333333334",
            "S08-avg": "-87.86$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.02$6.945416666666667"
        },
        {
            "date": "2020-02-01 02:00:00",
            "03-zsavg": "57.98$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "57.98$50.48291666666666",
            "02-zsavg": "6.06$3.3470833333333334",
            "S03-avg": "62.25$58.86958333333331",
            "01-zsavg": "2.30$1.6291666666666667",
            "S01-avg": "9.67$11.805000000000001",
            "S05-avg": "2.66$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "2.30$1.6291666666666667",
            "02-avg": "6.06$3.3470833333333334",
            "S08-avg": "-83.47$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.01$6.945416666666667"
        },
        {
            "date": "2020-02-01 03:00:00",
            "03-zsavg": "58.99$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "58.99$50.48291666666666",
            "02-zsavg": "5.19$3.3470833333333334",
            "S03-avg": "59.44$58.86958333333331",
            "01-zsavg": "2.31$1.6291666666666667",
            "S01-avg": "13.15$11.805000000000001",
            "S05-avg": "1.15$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "2.31$1.6291666666666667",
            "02-avg": "5.19$3.3470833333333334",
            "S08-avg": "-71.88$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.90$6.945416666666667"
        },
        {
            "date": "2020-02-01 04:00:00",
            "03-zsavg": "60.02$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "60.02$50.48291666666666",
            "02-zsavg": "1.97$3.3470833333333334",
            "S03-avg": "58.57$58.86958333333331",
            "01-zsavg": "2.47$1.6291666666666667",
            "S01-avg": "13.26$11.805000000000001",
            "S05-avg": "1.42$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "2.47$1.6291666666666667",
            "02-avg": "1.97$3.3470833333333334",
            "S08-avg": "-71.91$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.05$6.945416666666667"
        },
        {
            "date": "2020-02-01 05:00:00",
            "03-zsavg": "64.18$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "64.18$50.48291666666666",
            "02-zsavg": "3.63$3.3470833333333334",
            "S03-avg": "58.46$58.86958333333331",
            "01-zsavg": "2.73$1.6291666666666667",
            "S01-avg": "13.29$11.805000000000001",
            "S05-avg": "4.60$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "2.73$1.6291666666666667",
            "02-avg": "3.63$3.3470833333333334",
            "S08-avg": "-55.84$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.02$6.945416666666667"
        },
        {
            "date": "2020-02-01 06:00:00",
            "03-zsavg": "68.84$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "68.84$50.48291666666666",
            "02-zsavg": "3.62$3.3470833333333334",
            "S03-avg": "56.72$58.86958333333331",
            "01-zsavg": "3.25$1.6291666666666667",
            "S01-avg": "9.87$11.805000000000001",
            "S05-avg": "3.02$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "3.25$1.6291666666666667",
            "02-avg": "3.62$3.3470833333333334",
            "S08-avg": "-57.77$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.15$6.945416666666667"
        },
        {
            "date": "2020-02-01 07:00:00",
            "03-zsavg": "61.24$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "61.24$50.48291666666666",
            "02-zsavg": "2.57$3.3470833333333334",
            "S03-avg": "55.91$58.86958333333331",
            "01-zsavg": "3.12$1.6291666666666667",
            "S01-avg": "11.00$11.805000000000001",
            "S05-avg": "0.21$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "3.12$1.6291666666666667",
            "02-avg": "2.57$3.3470833333333334",
            "S08-avg": "-51.45$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.03$6.945416666666667"
        },
        {
            "date": "2020-02-01 08:00:00",
            "03-zsavg": "73.33$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "73.33$50.48291666666666",
            "02-zsavg": "3.58$3.3470833333333334",
            "S03-avg": "58.64$58.86958333333331",
            "01-zsavg": "3.09$1.6291666666666667",
            "S01-avg": "12.29$11.805000000000001",
            "S05-avg": "2.90$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "3.09$1.6291666666666667",
            "02-avg": "3.58$3.3470833333333334",
            "S08-avg": "-52.87$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.21$6.945416666666667"
        },
        {
            "date": "2020-02-01 09:00:00",
            "03-zsavg": "59.48$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "59.48$50.48291666666666",
            "02-zsavg": "3.47$3.3470833333333334",
            "S03-avg": "57.43$58.86958333333331",
            "01-zsavg": "2.20$1.6291666666666667",
            "S01-avg": "12.22$11.805000000000001",
            "S05-avg": "1.73$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "2.20$1.6291666666666667",
            "02-avg": "3.47$3.3470833333333334",
            "S08-avg": "-58.81$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.04$6.945416666666667"
        },
        {
            "date": "2020-02-01 10:00:00",
            "03-zsavg": "48.88$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "48.88$50.48291666666666",
            "02-zsavg": "3.73$3.3470833333333334",
            "S03-avg": "57.15$58.86958333333331",
            "01-zsavg": "1.29$1.6291666666666667",
            "S01-avg": "11.13$11.805000000000001",
            "S05-avg": "0.33$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "1.29$1.6291666666666667",
            "02-avg": "3.73$3.3470833333333334",
            "S08-avg": "-65.37$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.93$6.945416666666667"
        },
        {
            "date": "2020-02-01 11:00:00",
            "03-zsavg": "52.03$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "52.03$50.48291666666666",
            "02-zsavg": "1.35$3.3470833333333334",
            "S03-avg": "60.02$58.86958333333331",
            "01-zsavg": "0.99$1.6291666666666667",
            "S01-avg": "12.66$11.805000000000001",
            "S05-avg": "0.21$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "0.99$1.6291666666666667",
            "02-avg": "1.35$3.3470833333333334",
            "S08-avg": "-83.30$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.88$6.945416666666667"
        },
        {
            "date": "2020-02-01 12:00:00",
            "03-zsavg": "50.34$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "50.34$50.48291666666666",
            "02-zsavg": "0.96$3.3470833333333334",
            "S03-avg": "61.03$58.86958333333331",
            "01-zsavg": "0.82$1.6291666666666667",
            "S01-avg": "12.64$11.805000000000001",
            "S05-avg": "0.88$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "0.82$1.6291666666666667",
            "02-avg": "0.96$3.3470833333333334",
            "S08-avg": "-87.22$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.87$6.945416666666667"
        },
        {
            "date": "2020-02-01 13:00:00",
            "03-zsavg": "50.10$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "50.10$50.48291666666666",
            "02-zsavg": "1.45$3.3470833333333334",
            "S03-avg": "57.66$58.86958333333331",
            "01-zsavg": "0.89$1.6291666666666667",
            "S01-avg": "12.71$11.805000000000001",
            "S05-avg": "2.41$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "0.89$1.6291666666666667",
            "02-avg": "1.45$3.3470833333333334",
            "S08-avg": "-88.10$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.69$6.945416666666667"
        },
        {
            "date": "2020-02-01 14:00:00",
            "03-zsavg": "40.65$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "40.65$50.48291666666666",
            "02-zsavg": "2.66$3.3470833333333334",
            "S03-avg": "60.43$58.86958333333331",
            "01-zsavg": "0.95$1.6291666666666667",
            "S01-avg": "9.43$11.805000000000001",
            "S05-avg": "1.07$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "0.95$1.6291666666666667",
            "02-avg": "2.66$3.3470833333333334",
            "S08-avg": "-89.11$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.76$6.945416666666667"
        },
        {
            "date": "2020-02-01 15:00:00",
            "03-zsavg": "36.19$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "36.19$50.48291666666666",
            "02-zsavg": "3.02$3.3470833333333334",
            "S03-avg": "61.35$58.86958333333331",
            "01-zsavg": "0.95$1.6291666666666667",
            "S01-avg": "12.55$11.805000000000001",
            "S05-avg": "1.83$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "0.95$1.6291666666666667",
            "02-avg": "3.02$3.3470833333333334",
            "S08-avg": "-90.20$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.76$6.945416666666667"
        },
        {
            "date": "2020-02-01 16:00:00",
            "03-zsavg": "37.01$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "37.01$50.48291666666666",
            "02-zsavg": "2.13$3.3470833333333334",
            "S03-avg": "56.69$58.86958333333331",
            "01-zsavg": "0.97$1.6291666666666667",
            "S01-avg": "12.41$11.805000000000001",
            "S05-avg": "0.95$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "0.97$1.6291666666666667",
            "02-avg": "2.13$3.3470833333333334",
            "S08-avg": "-85.75$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.65$6.945416666666667"
        },
        {
            "date": "2020-02-01 17:00:00",
            "03-zsavg": "34.04$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "34.04$50.48291666666666",
            "02-zsavg": "4.19$3.3470833333333334",
            "S03-avg": "58.37$58.86958333333331",
            "01-zsavg": "1.00$1.6291666666666667",
            "S01-avg": "12.35$11.805000000000001",
            "S05-avg": "0.98$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "1.00$1.6291666666666667",
            "02-avg": "4.19$3.3470833333333334",
            "S08-avg": "-88.47$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.91$6.945416666666667"
        },
        {
            "date": "2020-02-01 18:00:00",
            "03-zsavg": "31.21$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "31.22$50.48291666666666",
            "02-zsavg": "4.82$3.3470833333333334",
            "S03-avg": "57.42$58.86958333333331",
            "01-zsavg": "1.05$1.6291666666666667",
            "S01-avg": "9.03$11.805000000000001",
            "S05-avg": "0.11$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "1.05$1.6291666666666667",
            "02-avg": "4.82$3.3470833333333334",
            "S08-avg": "-89.16$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.89$6.945416666666667"
        },
        {
            "date": "2020-02-01 19:00:00",
            "03-zsavg": "34.34$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "34.34$50.48291666666666",
            "02-zsavg": "4.47$3.3470833333333334",
            "S03-avg": "57.32$58.86958333333331",
            "01-zsavg": "0.83$1.6291666666666667",
            "S01-avg": "12.30$11.805000000000001",
            "S05-avg": "0.25$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "0.83$1.6291666666666667",
            "02-avg": "4.47$3.3470833333333334",
            "S08-avg": "-85.35$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.93$6.945416666666667"
        },
        {
            "date": "2020-02-01 20:00:00",
            "03-zsavg": "40.56$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "40.56$50.48291666666666",
            "02-zsavg": "1.16$3.3470833333333334",
            "S03-avg": "54.83$58.86958333333331",
            "01-zsavg": "0.71$1.6291666666666667",
            "S01-avg": "12.19$11.805000000000001",
            "S05-avg": "0.05$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "0.71$1.6291666666666667",
            "02-avg": "1.16$3.3470833333333334",
            "S08-avg": "-72.81$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.88$6.945416666666667"
        },
        {
            "date": "2020-02-01 21:00:00",
            "03-zsavg": "41.11$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "41.11$50.48291666666666",
            "02-zsavg": "3.78$3.3470833333333334",
            "S03-avg": "58.76$58.86958333333331",
            "01-zsavg": "0.72$1.6291666666666667",
            "S01-avg": "11.81$11.805000000000001",
            "S05-avg": "2.17$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "0.72$1.6291666666666667",
            "02-avg": "3.78$3.3470833333333334",
            "S08-avg": "-81.06$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.09$6.945416666666667"
        },
        {
            "date": "2020-02-01 22:00:00",
            "03-zsavg": "44.67$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "44.67$50.48291666666666",
            "02-zsavg": "6.50$3.3470833333333334",
            "S03-avg": "60.35$58.86958333333331",
            "01-zsavg": "0.58$1.6291666666666667",
            "S01-avg": "8.85$11.805000000000001",
            "S05-avg": "0.33$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "0.58$1.6291666666666667",
            "02-avg": "6.50$3.3470833333333334",
            "S08-avg": "-79.74$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "7.10$6.945416666666667"
        },
        {
            "date": "2020-02-01 23:00:00",
            "03-zsavg": "43.05$50.482499999999995",
            "S06-avg": "-$1.0",
            "03-avg": "43.05$50.48291666666666",
            "02-zsavg": "7.36$3.3470833333333334",
            "S03-avg": "58.35$58.86958333333331",
            "01-zsavg": "1.09$1.6291666666666667",
            "S01-avg": "11.91$11.805000000000001",
            "S05-avg": "0.16$1.493333333333333",
            "S07-avg": "-$1.0",
            "01-avg": "1.09$1.6291666666666667",
            "02-avg": "7.36$3.3470833333333334",
            "S08-avg": "-84.60$-77.08208333333333",
            "S04-avg": "-$1.0",
            "S02-avg": "6.82$6.945416666666667"
        }
    ],
    "recordsTotal": 1,
    "columns": [
        {
            "id": "01-avg",
            "name": "烟尘实测浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "01-zsavg",
            "name": "烟尘折算浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "02-avg",
            "name": "二氧化硫实测浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "02-zsavg",
            "name": "二氧化硫折算浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "03-avg",
            "name": "氮氧化物实测浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "03-zsavg",
            "name": "氮氧化物折算浓度(mg/m³)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "S01-avg",
            "name": "氧含量均值(%)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "S02-avg",
            "name": "烟气流速平均值(m/s)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "S03-avg",
            "name": "烟气温度平均值(℃)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "S04-avg",
            "name": "烟气动压平均值(兆帕)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "S05-avg",
            "name": "烟气湿度平均值(%)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "S06-avg",
            "name": "制冷温度平均值(℃)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "S07-avg",
            "name": "烟道截面积(㎡)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        },
        {
            "id": "S08-avg",
            "name": "烟气压力平均值(兆帕)",
            "type": null,
            "show": true,
            "render": null,
            "width": null,
            "exportWidth": null,
            "checked": true,
            "ordering": false,
            "children": null
        }
    ]
}
```