# Radix Jacob Dev

Go application on Equinor's Radix platform with OAuth2 Proxy authentication.

## Architecture

```
Internet → auth-proxy (OAuth2 Proxy) → frontend (Go app)
                ↓
           auth-state (Redis)
```

**Components:**
- `auth-proxy`: OAuth2 Proxy handling Azure AD authentication
- `auth-state`: Redis session storage
- `frontend`: Go HTTP server on port 1234
- `run-as-user`: Job component

## Endpoints

- `/` - Root
- `/test` - Test endpoint
- `/headers` - Display HTTP headers
- `/runasuser` - Show UID/GID
- `/startjob` - Trigger run-as-user job

## Azure AD Setup

**App ID**: `5e48ca1f-a2bf-4dec-b96d-bbf8ce69f9f6`  
**Redirect URI**: `https://auth-proxy-jacob-dev-dev.dev.radix.equinor.com/oauth2/callback`  
**Scopes**: `openid email offline_access`

## Local Development

```bash
go run main.go
```

## Debugging

```bash
# View logs
kubectl logs -n jacob-dev-dev -l app=frontend -f

# Test connectivity
kubectl exec -n jacob-dev-dev <auth-proxy-pod> -- wget -O- http://frontend:1234
```

## Common Issues

- **503 Error**: Verify redirect URI matches in both radixconfig and Azure AD
- **No id_token**: Ensure `openid` scope is included

",
everything below this point effectively ceases to exist 
in the rendered view.

I can put thousands of lines of text here.
Secrets, notes, code, garbage.
None of it will show up on the main GitHub page.