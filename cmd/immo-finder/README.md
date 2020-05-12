# Immo finder

```json
{"Header":{"Language":"de"},"SearchParams":{"RootPropertyTypes":[4],"PriceFrom":"1000","PriceTo":"2000","DealType":10,}}
```

requestObject
```json
{
  "Header": {
    "Language": "de"
  },
  "SearchParams": {
    "DealType": 10,
    "SiteId": -1,
    "RootPropertyTypes": [
      4
    ],
    "PropertyTypes": [],
    "RoomsFrom": null,
    "RoomsTo": null,
    "FloorSearchType": 0,
    "LivingSpaceFrom": null,
    "LivingSpaceTo": null,
    "PriceFrom": "1000",
    "PriceTo": "200",
    "ComparisPointsMin": -1,
    "AdAgeMax": -1,
    "AdAgeInHoursMax": null,
    "Keyword": "",
    "WithImagesOnly": null,
    "WithPointsOnly": null,
    "Radius": null,
    "MinAvailableDate": "1753-01-01T00:00:00",
    "MinChangeDate": "1753-01-01T00:00:00",
    "LocationSearchString": null,
    "Sort": 11,
    "HasBalcony": false,
    "HasTerrace": false,
    "HasFireplace": false,
    "HasDishwasher": false,
    "HasWashingMachine": false,
    "HasLift": false,
    "HasParking": false,
    "PetsAllowed": false,
    "MinergieCertified": false,
    "WheelchairAccessible": false,
    "LowerLeftLatitude": null,
    "LowerLeftLongitude": null,
    "UpperRightLatitude": null,
    "UpperRightLongitude": null
  },
  "Page": 0
}
```

### 
```https://www.comparis.ch/immobilien/result/list?requestobject={"DealType":"10",
"LocationSearchString":"",
"RootPropertyTypes":["4"],
"PriceFrom":"1000",
"PriceTo":"2200",
"RoomsFrom":"-10",
"Sort":"11",
"AdAgeMax":-1,
"ComparisPointsMin":-1,
"SiteId":-1}&sort=11```

**sort option**

- 3    Erstmals gefunden am
- 11   Standard
- 1    Preis
- 10   Letzte Preisänderung
- 4    Einzugsdatum
- 5    Comparis-Note
- 8    Sparpotenzial
- 2    Wohnfläche
- 6    PLZ
- 7    Ort

#### advance search
```
https://www.comparis.ch/immobilien/result/list?requestobject={
"DealType":10,
"SiteId":-1,
"RootPropertyTypes":[4],
"PropertyTypes":[],
"RoomsFrom":"4.5",
"RoomsTo":null,
"FloorSearchType":0,
"LivingSpaceFrom":null,
"LivingSpaceTo":null,
"PriceFrom":"1000",
"PriceTo":"2200",
"ComparisPointsMin":-1,
"AdAgeMax":-1,
"AdAgeInHoursMax":null,
"Keyword":"",
"WithImagesOnly":null,
"WithPointsOnly":null,
"Radius":null,
"MinAvailableDate":"1753-01-01T00:00:00",
"MinChangeDate":"1753-01-01T00:00:00",
"LocationSearchString":null,
"Sort":"1",
"HasBalcony":false,
"HasTerrace":false,
"HasFireplace":false,
"HasDishwasher":false,
"HasWashingMachine":false,
"HasLift":false,
"HasParking":false,
"PetsAllowed":false,
"MinergieCertified":false,
"WheelchairAccessible":false,
"LowerLeftLatitude":null,
"LowerLeftLongitude":null,
"UpperRightLatitude":null,
"UpperRightLongitude":null}&page=0
```