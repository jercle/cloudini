# Security and optimisation review

Review of the whole repository at v0.4.19. Part 1 lists what this branch
changes. Part 2 lists recommendations that were left alone because they need a
decision, a wider refactor, or an environment (PowerShell, Windows, live
tenants) that the changes could not be tested against.

## 1. Changed on this branch

### Security

| Area | Problem | Change |
| --- | --- | --- |
| Config and token cache encryption (`lib/fnCld.go`) | AES-CFB with a fixed, hard-coded IV and no integrity check. The same key and IV were reused for every write, so two versions of a file leak the XOR of their plaintexts, and a modified or wrongly-keyed file decrypted to garbage without an error. | New writes use AES-256-GCM with a random nonce (`cld2:` prefix). Tampering and wrong keys are now reported. Files in the old format are still read and are upgraded the next time they are saved. |
| File permissions | Config, key, token cache and decrypted exports were written with `os.ModePerm` (0777) or 0644; the config directory was created 0777. Any local user could read client secrets and bearer tokens. | Secrets are written 0600 and the config and cache directories are 0700, including tightening ones that already exist (`lib/fnSecureFiles.go`). Plain output files and directories use 0644 and 0755. |
| LDAP (`cmd/ad/getUsers.go`) | StartTLS ran with `InsecureSkipVerify: true`, so the bind password could be captured by anyone able to intercept the connection. | The domain controller certificate is verified by default, with TLS 1.2 as the minimum. A per-domain `insecureSkipVerify` config field re-enables the old behaviour and prints a warning. |
| Certificate auth (`cmd/azure/fnAuth.go`) | Key and certificate were read from a hard-coded `/home/jercle/.config/cld/` path, and a missing PEM block caused a nil dereference. | Paths come from the config directory and missing PEM data returns an error. |
| Error output (`cmd/azure/fnUtils.go`) | A failed GET serialised the whole `http.Response`, which includes the request and its `Authorization` header, into the error text. | The error contains only the URL and status. |
| Format strings | API response text was passed as the format argument of `fmt.Errorf` in about 20 places, which corrupts any message containing `%`. | Replaced with `errors.New` or explicit verbs. `lib.FormatError` now wraps the original error with `%w`. |
| Token request (`GetLogAnalyticsToken`) | The client secret was concatenated into a form body without encoding. | The body is built with `url.Values`. |
| Debug dumps (`lib/fnCertDumps.go`) | Wrote to a hard-coded home directory path. | Writes to the cache directory with 0600. |
| Repository hygiene | Nothing stopped a real config or key being committed. | `.gitignore` now covers `cldConf.json`, root-level key and certificate files, and build output. |

### Reliability and performance

| Area | Problem | Change |
| --- | --- | --- |
| HTTP (`lib/httpClient.go`) | Every call used `http.DefaultClient`: no timeout, so a stalled connection hung the CLI forever, and only two idle connections per host, so concurrent tenant goroutines kept re-doing TLS handshakes. | One shared client with a 5 minute timeout (`CLD_HTTP_TIMEOUT` overrides it, `0` disables it) and a larger idle pool. |
| Response bodies | About 25 call sites never closed `res.Body`, leaking a connection each and preventing keep-alive reuse. | Every call site closes the body. |
| Azure token cache | Tokens were saved under one key and looked up under another, so the cache never hit and every command requested a new token per tenant and scope. Read and write tokens also shared a key. | Save and lookup use the same key. Tokens within 2 minutes of expiry count as expired. |
| Token cache file | Concurrent goroutines did unsynchronised read-modify-write on the cache file, a corrupt cache was fatal, and file handles were leaked. | Access is serialised by a mutex, and an unreadable cache is treated as empty. |
| First run (`InitConfig`) | With no config file present the CLI crashed with a nil pointer dereference. | A default config is created at the path that was asked for. |
| `GetServicePrincipalToken` | With `NoCache` set and a mutex supplied, the mutex was never released. | The lock is only taken around the cache read. |
| `GetToken` | Nil dereference when no token cache existed. | Nil check added. |
| SharePoint upload | The response was ignored, so a failed upload was reported as uploaded. | Non-2xx responses are fatal and include the response text. |
| `GetWebsiteCertificate` | Panicked with a nil dereference for non-TLS URLs. | Explicit message. |
| Struct tag (`ListGalleryImageVersionsResponse.NextLink`) | Malformed tag, so the `bson` name was ignored. | Fixed. |
| Build | The Makefile did not set `GOEXPERIMENT=jsonv2`, so `make` failed, and binaries carried debug tables and local paths. | Makefile sets the experiment and uses `-trimpath -s -w`. GoReleaser uses `-trimpath`. |
| CI | No build, vet or test on pull requests. | `.github/workflows/ci.yml` builds for all three operating systems, vets, runs the `lib` tests and runs `govulncheck`. |

`go vet` previously failed in `lib`, which also stopped `go test ./lib/` from
running. It now passes for every package that loads on Linux, and
`lib/fnCld_test.go` covers the encryption formats, the token cache and file
permissions.

### Behaviour changes to be aware of

- **LDAP certificate verification is on.** If a domain controller uses a
  certificate the local machine does not trust, add `"insecureSkipVerify": true`
  to that domain in the config, or better, trust the issuing CA.
- **Cached Azure tokens are now actually reused** until shortly before expiry.
  After changing a service principal's roles, run `cld config -c` to clear the
  cache.
- **Requests time out after 5 minutes.** Set `CLD_HTTP_TIMEOUT` for very large
  uploads or slow links.
- **Existing encrypted configs keep working.** Re-import or save the config once
  to move it to the new format. The old format is not removed.

## 2. Recommended, not changed

### Security

1. **Access token on the PowerShell command line** (`cmd/m365/funcExchange.go`).
   The Exchange Online token is passed as an argument to `pwsh`, so it is
   visible to other local users in the process list. Pass it through an
   environment variable or stdin and read it in the `.ps1` scripts.
2. **Key handling for config encryption.** `CLD_CONFIG_SECRET` is used directly
   as the AES key, so it must be exactly 32 characters and has only the entropy
   of what was typed. Derive the key with Argon2id or scrypt and a stored salt.
   Without `CLD_CONFIG_SECRET`, the generated key sits in `cache/key` beside the
   data it protects, which only guards against casual reading; the OS keychain
   (Keychain, DPAPI, libsecret) would be the stronger place for it.
3. **Secrets in the config file.** Client secrets, the LDAP bind password and
   the MongoDB connection string are stored in plain JSON unless
   `CLD_CONFIG_ENCRYPT` is set. Consider making encryption the default, or
   moving to certificate or workload identity credentials. The Forgerock
   `clientSecretBase64` field is encoding, not protection.
4. **`--export` and `--exportToken`** write decrypted secrets to a path of the
   user's choosing. They are now 0600, but a printed warning would help.
5. **Release signing** (`.goreleaser.yml`) reads the GPG passphrase from a
   plaintext `./gpg_pwd` file. Use the agent or a CI secret.
6. **Dependencies.** `govulncheck` could not be run in the review environment
   because of the Go 1.27 toolchain requirement; the new CI job runs it. Several
   direct dependencies are old or unmaintained and worth replacing:
   `StackExchange/wmi`, `github/smimesign`, `TylerBrock/colorjson`,
   `pkg/errors`, `rmasci/ipsubnet`. `go.mongodb.org/mongo-driver` v1 is in
   maintenance; v2 is current.
7. **Pin GitHub Actions to commit SHAs** rather than moving tags, and move the
   CodeQL job to `ubuntu-latest` unless Windows-only code needs analysing.
8. **Query construction.** Resource Graph and Log Analytics request bodies are
   built by string concatenation (`fnLogAnalytics.go:209`, `fnUtils.go`).
   Marshal a struct instead so quotes in input cannot break the JSON.
9. **Vendored `lib/certstore` and `lib/certtostorelocal`** are copies of
   upstream projects and will not receive upstream fixes. `go vet` reports
   `unsafe.Pointer` misuse in `certtostore_windows.go`. Track upstream or record
   the commit they were copied from.

### Performance

1. **Config is re-read on every call.** `GetCldConfig` has about 60 call sites
   and `InitConfig` runs on every token cache access; each one stats, reads and
   possibly decrypts the file, or calls Azure App Configuration over the
   network. Load once per process with `sync.Once`.
2. **Unbounded fan-out.** Around 30 loops start one goroutine per tenant,
   subscription or resource with no limit, which invites ARM throttling (429).
   Use `errgroup.SetLimit` and honour `Retry-After`.
3. **No retries.** A single 429 or 5xx fails the command. A small retry with
   backoff in the shared HTTP helpers would cover every caller.
4. **No cancellation.** Requests use `context.Background()` or no context, so
   Ctrl-C cannot stop in-flight work cleanly. Thread the cobra command context
   through the HTTP helpers.
5. **Whole files in memory.** Packer logs, cost exports and SharePoint uploads
   are read fully into memory. Stream with `bufio.Scanner` and `os.Open`.
6. **JSON round-trips for conversion.** Several places marshal a value and
   unmarshal it into another type (`cmd/ad/funcs.go`, `GetAzureResourceTypes`,
   `GetWebsiteCertificate`). Direct field mapping is faster and type checked.
7. **Token cache format.** Each token is JSON, then base64, inside JSON, then
   encrypted and base64 again. Storing the tokens directly in the map removes
   two encoding passes.

### Code health

1. **`log.Fatal` in library code.** `lib.CheckFatalError` is called several
   hundred times, including inside goroutines, where it exits without running
   deferred cleanup such as `Disconnect` or stopping spinners. Return errors and
   exit in the cobra `RunE` handlers.
2. **`go build ./...` and `go vet ./...` fail.** `scripts/` holds an empty
   `generateAzureResourceList.go` and two `main` functions in one package.
   Move each script to its own directory or add `//go:build ignore`.
3. **Existing tests fail.** `cmd/utils/datatransforms` panics in an example test
   and `lib/certtostorelocal` needs testdata that is git-ignored.
4. **`cmd/ad/getUsers.go`** prints the search result and calls `os.Exit(0)`, so
   the code after it never runs.
5. **`GetEXOMailboxesWithPermissionsForTenant`** checks the stderr buffer before
   the process has started, so that check can never fire.
6. **Large blocks of commented-out code** make review harder; git history
   already preserves them.
