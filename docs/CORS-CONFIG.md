# CORS Configuration

This document explains how to configure Cross-Origin Resource Sharing (CORS) in this application.

## Overview

Cross-Origin Resource Sharing (CORS) is a mechanism that allows resources on a web page to be requested from a domain outside the domain from which the resource originated. Our application uses the [go-chi/cors](https://github.com/go-chi/cors) package to implement CORS.

## Configuration

CORS can be configured through environment variables. Add these variables to your `.env` file:

| Environment Variable | Description | Default Value |
|---------------------|-------------|---------------|
| `CORS_ALLOWED_ORIGINS` | Comma-separated list of allowed origins | `*` (all origins) |
| `CORS_ALLOWED_METHODS` | Comma-separated list of allowed HTTP methods | `GET,POST,PUT,DELETE,OPTIONS` |
| `CORS_ALLOWED_HEADERS` | Comma-separated list of allowed headers | `Accept,Authorization,Content-Type,X-CSRF-Token` |
| `CORS_EXPOSED_HEADERS` | Comma-separated list of headers that browsers are allowed to access | `Link` |
| `CORS_ALLOW_CREDENTIALS` | Allow cookies and credentials | `true` |
| `CORS_MAX_AGE` | How long the results of a preflight request can be cached (in seconds) | `300` |
| `CORS_OPTIONS_PASSTHROUGH` | Whether to pass OPTIONS requests to handlers | `false` |
| `CORS_DEBUG` | Enable debug mode to log CORS-related messages | `false` |

## Examples

### Allow all origins (default)

```
CORS_ALLOWED_ORIGINS=*
```

### Allow specific origins

```
CORS_ALLOWED_ORIGINS=https://myapp.com,https://admin.myapp.com
```

### Allow wildcard subdomains

```
CORS_ALLOWED_ORIGINS=https://*.myapp.com
```

### Allow both HTTP and HTTPS

```
CORS_ALLOWED_ORIGINS=https://*,http://*
```

## Important Notes

- If `CORS_ALLOWED_ORIGINS` contains `*`, all origins will be allowed.
- If `CORS_ALLOWED_HEADERS` contains `*`, all headers will be allowed.
- Setting `CORS_ALLOW_CREDENTIALS` to `true` means that cookies will be included in cross-origin requests.
- A `CORS_MAX_AGE` of 300 seconds (5 minutes) is the maximum value not ignored by any major browsers.
- Setting `CORS_DEBUG` to `true` will log additional information about CORS-related issues.

## Security Considerations

- Use specific origins rather than `*` in production environments.
- Only allow the methods and headers that your application actually needs.
- Be cautious with `CORS_ALLOW_CREDENTIALS=true` as it can lead to security issues if not properly handled. 