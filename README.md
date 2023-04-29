# tracehound

Trace scanning and hopefully a notification tool

Goal: tracehound will do one or more searches on traces it can find, and barks when a search fails based on predefined search definitions

```
searches:
  - serviceName: shop-backend
    allowedMethods:
      - GET
      - POST
    bark: true
    maxLatencyMs: 1000
```
TraceQL is currently used in testing
