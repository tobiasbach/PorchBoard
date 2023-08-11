# Viessmann API

Free basic package: 1450 API calls per day, 1 active client, limited feature availability  

[Authentication](https://documentation.viessmann.com/static/authentication)  
\- how to retrieve access tokens  
\- be aware of broad TTL range from 20seconds to 180days

[Personal Dashboard](https://app.developer.viessmann.com/)  
\- manage OAuth client & personal data

[Data Points](https://documentation.viessmann.com/static/iot/data-points)  
\- overview of available data points depending on subscription



## Notes from first auth attempt

Switch from corporate to personal Postman account

There is a [public collection](https://www.postman.com/vimicho/workspace/viessmann-api-public/collection/12055031-17157e90-a2e8-47b6-a7b8-2320c2941db3?action=share&creator=12055031) with prepared requests available for Postman

1st request in browser 'Authorization request'
  requires code challenge
  scope param: 'IoTUser offline_access'

  returns code TTL 20 seconds!

2nd request via Postman 'Authorization code exchange'
  requires response code from step one
  requires code verifier

  returns Bearer token TTL 3600 seconds
  returns refresh_token (if scope in step one includes 'offline_access') TTL 180 days / 15552000 seconds

3rd request 'Authorization code exchange'
  requires refresh token

## Notes from first data request

Need to pass installation ID, gateway serial and device ID with the request. Don't forget the Bearer token for Authorization.
https://api.viessmann.com/iot/v2/features/installations/INSTALLATION_ID/gateways/GATEWAY_SERIAL/devices/DEVICE_ID/features/heating.sensors.temperature.outside

Request successful, response body in expected format with expected data

WIP: Unmarshaling / accessing data from the response body returns unexpected value:

RESPONSE BODY:
{"data":{"properties":{"value":{"type":"number","value":18.1,"unit":"celsius"},"status":{"type":"string","value":"connected"}},"commands":{},"apiVersion":1,"uri":"https://api.viessmann.com/iot/v2/features/installations/xxx/gateways/xxx/devices/xxx/features/heating.sensors.temperature.outside","gatewayId":"xxx","feature":"heating.sensors.temperature.outside","timestamp":"2023-08-11T20:47:58.613Z","isEnabled":true,"isReady":true,"deviceId":"0"}}

UNMARSHALED OUTSIDE TEMPERATURE COMPLETE:
{{{1 {} 0 heating.sensors.temperature.outside xxx true true {{string connected} {number celsius 0}} 2023-08-11T20:47:58.613Z https://api.viessmann.com/iot/v2/features/installations/xxx/gateways/xxx/devices/xxx/features/heating.sensors.temperature.outside}}}
