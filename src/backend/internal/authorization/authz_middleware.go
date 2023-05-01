package authorization

import "github.com/casbin/casbin/v2"

a := gormadapter.NewAdapterByDB(db)

e, err := casbin.NewEnforcer("path/to/model.conf", a)
if err != nil {
    log.Fatalf("failed to create casbin enforcer: %v", err)
}

// Allow unauthenticated users to access redirects, login, and register
e.AddPolicy("unauthenticated", "/r/*", "GET")
e.AddPolicy("unauthenticated", "/auth/register", "POST")
e.AddPolicy("unauthenticated", "/auth/login", "POST")

// Allow authenticated users to access redirects and any object they own
e.AddPolicy("authenticated", "/api/v1/url/*", "GET,PUT,POST,DELETE")
e.AddPolicy("authenticated", "/api/objects/*", "GET,POST,PUT,DELETE")
e.AddPolicy("authenticated", "/api/objects/*", "owner = sub")

// Allow admins to do anything
e.AddPolicy("admin", "*", "*")

// Check authorization before handling each request
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        sub := r.Header.Get("X-User")
        obj := r.URL.Path
        act := r.Method

        allowed, err := e.Enforce(sub, obj, act)
        if err != nil {
            log.Printf("failed to enforce authorization: %v", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        if !allowed {
            w.WriteHeader(http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}

// Use the authMiddleware to protect all routes
r := mux.NewRouter()
r.Use(authMiddleware)

